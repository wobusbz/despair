package app

import (
	"fmt"
	"os"
	"path/filepath"
)

var (
	App = &app{}
)

func init() {
	App.Name = "wobusbz"
	App.Version = "0.3"
	App.RootDir = interRootDir()
}

type app struct {
	Name    string // 项目名称
	Version string // 项目版本

	RootDir string // 项目根目录
}

func interRootDir() string {
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println("interRootDir error ", err)
		return " "
	}
	return filepath.Dir(cwd)
}
