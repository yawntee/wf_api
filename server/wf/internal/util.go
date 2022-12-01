package internal

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"github.com/mholt/archiver"
	"github.com/pkg/errors"
	"github.com/shamaton/msgpack/v2"
	"golang.org/x/exp/slog"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync/atomic"
	"time"
	"wf_api/server/wf/internal/asset"
	"wf_api/server/wf/internal/context"
)

func EncodeHeader(header string) string {
	var builder strings.Builder
	for _, ch := range header {
		switch ch {
		case '%':
			builder.WriteString("%25")
		case ' ':
			builder.WriteString("%20")
		case '+':
			builder.WriteString("%2B")
		case '/':
			builder.WriteString("%2F")
		case '?':
			builder.WriteString("%3F")
		case '#':
			builder.WriteString("%23")
		case '&':
			builder.WriteString("%26")
		default:
			builder.WriteByte(byte(ch))
		}

	}
	return builder.String()
}

func WrapMsgpack(data any) []byte {
	marshal, err := msgpack.Marshal(data)
	if err != nil {
		return nil
	}
	return []byte(base64.StdEncoding.EncodeToString(marshal))
}

func UnwrapMsgpack(reader *io.ReadCloser, v any) error {
	var buf bytes.Buffer
	_, err := buf.ReadFrom(base64.NewDecoder(base64.StdEncoding, *reader))
	if err != nil {
		return err
	}
	err = msgpack.Unmarshal(buf.Bytes(), v)
	if err != nil {
		return err
	}
	return nil
}

const (
	assetCacheDir = "asset_cache"
	retryLimit    = 3
)

func StartUpdateAssets(data GameUpdateData) {
	if atomic.AddInt32(&context.UpdateMutex, 1) > 1 {
		atomic.AddInt32(&context.UpdateMutex, -1)
		panic(context.ErrAssetUpdate)
	}
	go updateAssets(data)
}

func updateAssets(data GameUpdateData) {
	//重试次数
mainloop:
	for i := 1; i <= retryLimit; i++ {
		time.Sleep(2 * time.Second)
		//创建缓存文件夹
		stat, err := os.Stat(assetCacheDir)
		if stat != nil && stat.IsDir() {
			//存在则删除缓存文件
			dir, err := os.ReadDir(assetCacheDir)
			if err != nil {
				if i == retryLimit {
					assertUpdateFail(err)
				}
				fmt.Printf("更新出错，2秒后重试。。。\n%+v\n", errors.WithStack(err))
				continue mainloop
			}
			for _, entry := range dir {
				if !entry.IsDir() {
					_ = os.Remove(filepath.Join(assetCacheDir, entry.Name()))
				}
			}
			slog.Info("缓存文件夹清理完毕")
		} else {
			//不存在则创建
			err := os.Mkdir(assetCacheDir, os.ModePerm)
			if err != nil {
				if i == retryLimit {
					assertUpdateFail(err)
				}
				fmt.Printf("更新出错，2秒后重试。。。\n%+v\n", errors.WithStack(err))
				continue mainloop
			}
			slog.Info("缓存文件夹创建完毕")
		}
		//下载更新资源包
		for _, ver := range data.Diff {
			for _, packInfo := range ver.Archive {
				pack, err := http.Get(packInfo.Location)
				if err != nil {
					if i == retryLimit {
						assertUpdateFail(err)
					}
					fmt.Printf("更新出错，2秒后重试。。。\n%+v\n", errors.WithStack(err))
					continue mainloop
				}
				out, err := os.Create(filepath.Join(assetCacheDir, packInfo.Location[strings.LastIndex(packInfo.Location, "/")+1:]))
				if err != nil {
					if i == retryLimit {
						assertUpdateFail(err)
					}
					fmt.Printf("更新出错，2秒后重试。。。\n%+v\n", errors.WithStack(err))
					continue mainloop
				}
				size, err := io.Copy(out, pack.Body)
				if err != nil {
					if i == retryLimit {
						assertUpdateFail(err)
					}
					fmt.Printf("更新出错，2秒后重试。。。\n%+v\n", errors.WithStack(err))
					continue mainloop
				}
				_ = out.Close()
				slog.Info("资源包已下载", "StringId", out.Name(), "Size", size)
			}
		}
		//解压资源包
		packs, err := os.ReadDir(assetCacheDir)
		if err != nil {
			if i == retryLimit {
				assertUpdateFail(err)
			}
			fmt.Printf("更新出错，2秒后重试。。。\n%+v\n", errors.WithStack(err))
			continue mainloop
		}
		zipper := archiver.NewZip()
		zipper.OverwriteExisting = true
		for _, pack := range packs {
			if !pack.IsDir() && strings.HasSuffix(pack.Name(), ".zip") {
				slog.Info("解压资源包", "StringId", pack.Name())
				err := zipper.Unarchive(filepath.Join(assetCacheDir, pack.Name()), asset.DownloadAssetPath)
				if err != nil {
					if i == retryLimit {
						assertUpdateFail(err)
					}
					fmt.Printf("更新出错，2秒后重试。。。\n%+v\n", errors.WithStack(err))
					continue mainloop
				}
			}
		}
		//重置缓存
		asset.GlobalAsset.Reset()
		slog.Info("已重置资源缓存")
		//修改配置文件
		GlobalConfig.ResVer = data.Info.EventualTargetAssetVersion
		err = GlobalConfig.Flush()
		if err != nil {
			if i == retryLimit {
				assertUpdateFail(err)
			}
			fmt.Printf("更新出错，2秒后重试。。。\n%+v\n", errors.WithStack(err))
			continue mainloop
		}
		slog.Info("更新完毕", "ResVer", data.Info.TargetAssetVersion)
		//解锁
		atomic.AddInt32(&context.UpdateMutex, -1)
		break mainloop
	}
}

func assertUpdateFail(err error) {
	if err != nil {
		DebugTitleMsg("更新资源失败", errors.WithStack(err))
		os.Exit(1)
	}
}
