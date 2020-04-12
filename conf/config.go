package conf

import (
	"despair/app"
	"flag"
	"fmt"
	"github.com/BurntSushi/toml"
	"sync"
)

var (
	config = flag.String("-f", fmt.Sprintf("%s/conf/config.toml", app.App.RootDir), "conf default dir")
)

type TomlConfing struct {
	Title string
	Http  http
	Log   log
	Dbs   dbs
}

type http struct {
	Host string `toml:"host"`
}

type log struct {
	LogDir string `toml:"logDir"`
}

type dbs struct {
	Mysql string `toml:"mysql"`
	Redis string `redis:"redis"`
}

var (
	Conf *TomlConfing
	once sync.Once
)

// 读取配置文件
func InitConfig() {
	once.Do(func() {
		flag.Parse()
		if _, err := toml.DecodeFile(*config, &Conf); err != nil {
			return
		}
	})
}
