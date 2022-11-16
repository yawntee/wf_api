package internal

import (
	"fmt"
	"golang.org/x/exp/slog"
	"gopkg.in/yaml.v3"
	"os"
)

var Config struct {
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

func init() {
	file, err := os.Open("config.yaml")
	if err != nil {
		panic(fmt.Errorf("配置文件打开失败\n%w", err))
	}
	err = yaml.NewDecoder(file).Decode(&Config)
	if err != nil {
		panic(fmt.Errorf("未能正确加载配置文件\n%w", err))
	}
	if Config.Debug {
		slog.Default().Enabled(slog.DebugLevel)
	}
}
