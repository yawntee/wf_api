package asset

import (
	"fmt"
	"github.com/pkg/errors"
	"golang.org/x/exp/maps"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var (
	RootPath         = "./resources"
	AssetDownloadDir = filepath.Join(RootPath, "asset_download/dummy/download")
	ResourceDirs     = []string{
		AssetDownloadDir,
	}
)

var GlobalAsset = NewAsset()

type Asset struct {
	Cache            map[string]any
	refreshListeners []func()
}

func NewAsset() *Asset {
	asset := Asset{
		Cache: make(map[string]any),
	}
	return &asset
}

func (a *Asset) resolve(path string) string {
	hash := hashPath(path)
	for _, dir := range ResourceDirs {
		format := "%s"
		switch {
		case strings.HasSuffix(path, ".png"):
			format = "android_%s"
		}
		path = filepath.Join(dir, "production", fmt.Sprintf(format, "upload"), hash[:2], hash[2:])
		_, err := os.Stat(path)
		if errors.Is(err, os.ErrNotExist) {
			continue
		}
		if err != nil {
			panic(err)
		}
		return path
	}
	log.Printf("Cannot resolve file %s", path)
	return ""
}

func (a *Asset) Reset() {
	maps.Clear(a.Cache)
	for _, listener := range a.refreshListeners {
		listener()
	}
}

func (a *Asset) AddOnResetListener(callback func()) {
	a.refreshListeners = append(a.refreshListeners, callback)
}

func (a *Asset) GetTableFile(path string) *os.File {
	path = a.resolve("master" + path + ".orderedmap")
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
	path = a.resolve(path + ".atlas.amf3.deflate")
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
	path = a.resolve(path + ".png")
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
