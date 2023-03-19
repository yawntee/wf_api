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
	"time"
	"wf_api/wf/internal/asset"
	"wf_api/wf/internal/context"
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
	if context.UpdateMutex.TryLock() {
		go updateAssets(data)
	}
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
		zipper := archiver.NewZip()
		zipper.OverwriteExisting = true
		packs := data.Full.Archive

		for _, ver := range data.Diff {
			for _, packInfo := range ver.Archive {
				packs = append(packs, packInfo)
			}
		}

		for _, packInfo := range packs {
			pack, err := http.Get(packInfo.Location)
			if err != nil {
				if i == retryLimit {
					assertUpdateFail(err)
				}
				fmt.Printf("更新出错，2秒后重试。。。\n%+v\n", errors.WithStack(err))
				continue mainloop
			}
			filename := packInfo.Location[strings.LastIndex(packInfo.Location, "/")+1:]
			out, err := os.OpenFile(filepath.Join(assetCacheDir, filename), os.O_RDWR|os.O_CREATE, os.ModePerm)
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
			//解压资源包
			slog.Info("解压资源包", "StringId", filename)
			err = zipper.Unarchive(filepath.Join(assetCacheDir, filename), asset.AssetDownloadDir)
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
		slog.Info("更新完毕", "ResVer", data.Info.EventualTargetAssetVersion)
		//解锁
		context.UpdateMutex.Unlock()
		//清理缓存
		err = os.RemoveAll(assetCacheDir)
		if err != nil {
			slog.Info("缓存文件夹清理失败", err)
		} else {
			slog.Info("缓存文件夹清理完毕")
		}
		break mainloop
	}
}

func assertUpdateFail(err error) {
	if err != nil {
		DebugTitleMsg("更新资源失败", errors.WithStack(err))
		os.Exit(1)
	}
}
