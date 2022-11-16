package asset

import (
	"io"
	"os"
)

type CacheAsset struct {
	Asset map[string]string
	Cache map[string]any
}

func NewCacheAsset(paths ...string) *CacheAsset {
	asset := CacheAsset{
		Asset: make(map[string]string),
		Cache: make(map[string]any),
	}
	for _, path := range paths {
		dirs, err := os.ReadDir(path)
		if err != nil {
			panic(err)
		}
		for _, dir := range dirs {
			sha1Prefixes, err := os.ReadDir(path + "/" + dir.Name())
			if err != nil {
				panic(err)
			}
			for _, prefix := range sha1Prefixes {
				sha1Postfix, err := os.ReadDir(path + "/" + dir.Name() + "/" + prefix.Name())
				if err != nil {
					panic(err)
				}
				for _, postfix := range sha1Postfix {
					asset.Asset[prefix.Name()+postfix.Name()] = path + "/" + dir.Name() + "/" + prefix.Name() + "/" + postfix.Name()
				}
			}
		}
	}
	return &asset
}

func (a *CacheAsset) getTableFile(path string) io.Reader {
	file, err := os.Open(a.Asset[hashPath("master"+path+".orderedmap")])
	if err != nil {
		panic(err)
	}
	return file
}
