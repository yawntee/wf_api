package internal

import (
	"fmt"
	"golang.org/x/exp/slog"
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
)

var (
	ConfigFilePath = filepath.Join(os.Getenv("WF_DIR"), "config.yaml")
)

type Config struct {
	Debug       bool   `json:"debug" yaml:"debug"`
	OsVer       string `json:"osVer" yaml:"osVer"`
	VersionName string `json:"versionName" yaml:"versionName"`
	VersionCode int    `json:"versionCode" yaml:"versionCode"`
	DeviceName  string `json:"deviceName" yaml:"deviceName"`
	Game        string `json:"game" yaml:"game"`
	Os          string `json:"os" yaml:"os"`
	Mmid        string `json:"mmid" yaml:"mmid"`
	CheckAuth   string `json:"checkAuth" yaml:"checkAuth"`
	Accompany   string `json:"accompany" yaml:"accompany"`
	Face        string `json:"face" yaml:"face"`
	ResVer      string `json:"resVer" yaml:"resVer"`
	LtKid       string `json:"LtKid" yaml:"LtKid"`
}

var GlobalConfig Config

func (c *Config) Flush() error {
	marshal, err := yaml.Marshal(c)
	if err != nil {
		return err
	}
	err = os.WriteFile(ConfigFilePath, marshal, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

func init() {
	file, err := os.Open(ConfigFilePath)
	if err != nil {
		panic(fmt.Errorf("配置文件打开失败\n%w", err))
	}
	err = yaml.NewDecoder(file).Decode(&GlobalConfig)
	if err != nil {
		panic(fmt.Errorf("未能正确加载配置文件\n%w", err))
	}
	if GlobalConfig.Debug {
		logger := slog.New(slog.HandlerOptions{Level: slog.DebugLevel}.NewTextHandler(os.Stderr))
		slog.SetDefault(logger)
	}
}
