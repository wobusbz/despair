package main

import (
	"despair/app/routers"
	"despair/common/logger"
	"despair/conf"
	"despair/db"
	"github.com/gin-gonic/gin"
)

func init() {
	conf.InitConfig()
	logger.InitLogger()
	db.InitMysql("mysql", "root:@tcp(127.0.0.1:3306)/wuhuarou?charset=utf8&parseTime=True&loc=Local")
	logger.Debug("mysql connect succ")
}

func main() {

	r := gin.Default()
	// 中间件

	// 注册前台路由
	indexRouter := r.Group("")
	routers.ResistIndexRoutes(indexRouter)
	logger.Debug("abc%s", "abc")
	// 注册后端路由
	adminRouter := r.Group("/admin")
	routers.RegisterAdminRouter(adminRouter)
	r.Run("0.0.0.0:8800")
}
