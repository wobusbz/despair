package conf

import (
	"flag"
	"sync"
)

var (
	once   sync.Once
	config = flag.String("-f", "common/conf", "conf default dir")
)

func Init() {
	once.Do(func() {

	})
}

// mysql dbhttps://github.com/spf13/viper
var (
	Mysql_Host     = "127.0.0.1"
	Mysql_Port     = 3306
	Mysql_User     = "root"
	Mysql_Pass     = ""
	Mtsql_Database = ""
	Mysql_Charset  = "utf8"
)

// log
var (
	Log_Path = ""
)

// app addrs
var (
	ADDRS = "127.0.0.1"
	PORT  = "8080"
)
