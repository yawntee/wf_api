package asset

import (
	"fmt"
	"golang.org/x/exp/maps"
	"os"
	"path/filepath"
)

var (
	RootPath          = filepath.Join(os.Getenv("WF_DIR"), "resources")
	DownloadAssetPath = filepath.Join(RootPath, "asset_download/dummy/download")
)

var GlobalAsset = NewAsset()

type Asset struct {
	Mapper           map[string]string
	Cache            map[string]any
	refreshListeners []func()
}

func NewAsset() *Asset {
	asset := Asset{
		Mapper: make(map[string]string),
		Cache:  make(map[string]any),
	}
	asset.init()
	return &asset
}

func (a *Asset) init() {
	//a.resolve(BundleAssetPath)
	a.resolve(DownloadAssetPath)
}

func (a *Asset) resolve(path string) {
	path = filepath.Join(path, "production")
	dirs, err := os.ReadDir(path)
	if err != nil {
		panic(err)
	}
	for _, dir := range dirs {
		if dir.IsDir() {
			sha1Prefixes, err := os.ReadDir(filepath.Join(path, dir.Name()))
			if err != nil {
				panic(err)
			}
			for _, prefix := range sha1Prefixes {
				sha1Postfix, err := os.ReadDir(filepath.Join(path, dir.Name(), prefix.Name()))
				if err != nil {
					panic(err)
				}
				for _, postfix := range sha1Postfix {
					hash := prefix.Name() + postfix.Name()
					if _, ok := a.Mapper[hash]; ok {
						panic(fmt.Errorf("HASH冲突：%s", hash))
					}
					a.Mapper[hash] = filepath.Join(path, dir.Name(), prefix.Name(), postfix.Name())
				}
			}
		}

	}
}

func (a *Asset) Reset() {
	maps.Clear(a.Mapper)
	maps.Clear(a.Cache)
	a.init()
	for _, listener := range a.refreshListeners {
		listener()
	}
}

func (a *Asset) AddOnResetListener(callback func()) {
	a.refreshListeners = append(a.refreshListeners, callback)
}

func (a *Asset) GetTableFile(path string) *os.File {
	path = a.getRealPath("master" + path + ".orderedmap")
	fmt.Println(path)
	if path == "" {
		return nil
	}
	file, err := os.Open(path)
	if err != nil {
		return nil
	}
	return file
}

func (a *Asset) GetSpriteSheet(path string) *os.File {
	path = a.getRealPath(path + ".atlas.amf3.deflate")
	if path == "" {
		return nil
	}
	file, err := os.Open(path)
	if err != nil {
		return nil
	}
	return file
}

func (a *Asset) GetPicture(path string) []byte {
	path = a.getRealPath(path + ".png")
	if path == "" {
		return nil
	}
	pic, err := os.ReadFile(path)
	if err != nil {
		return nil
	}
	pic[1] = 0x50
	pic[2] = 0x4E
	pic[3] = 0x47
	return pic
}

func (a *Asset) getRealPath(path string) string {
	if hashPath, ok := a.Mapper[hashPath(path)]; ok {
		return hashPath
	} else {
		fmt.Printf("warning: 无法找到资源文件：%s\n", path)
		return ""
	}
}
