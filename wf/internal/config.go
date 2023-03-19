package internal

import (
	"github.com/pkg/errors"
	"golang.org/x/exp/slog"
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"path/filepath"
)

var (
	DataDir        = "data"
	ConfigFilePath = filepath.Join(DataDir, "config.yaml")
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

var GlobalConfig = Config{
	Debug:       false,
	OsVer:       "12",
	VersionName: "1.6.3",
	VersionCode: 1006003,
	DeviceName:  "M1912G7BE",
	Game:        "wf",
	Os:          "1",
	Mmid:        "",
	CheckAuth:   "1",
	Accompany:   "1",
	Face:        "1",
	ResVer:      "",
	LtKid:       "v3.0.14.2",
}

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
		if errors.Is(err, os.ErrNotExist) {
			_ = GlobalConfig.Flush()
			return
		}
		log.Printf("配置文件打开失败\n%w", err)
		return
	}
	err = yaml.NewDecoder(file).Decode(&GlobalConfig)
	if err != nil {
		log.Printf("未能正确加载配置文件\n%w", err)
		return
	}
	if GlobalConfig.Debug {
		logger := slog.New(slog.HandlerOptions{Level: slog.DebugLevel}.NewTextHandler(os.Stderr))
		slog.SetDefault(logger)
	}
}
