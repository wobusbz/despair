package main

import (
	"despair/Db"
	"despair/app/routers"
	"github.com/gin-gonic/gin"
)

func init() {
	Db.InitMysql("mysql", "root:@tcp(127.0.0.1:3306)/wuhuarou?charset=utf8&parseTime=True&loc=Local")
}

func main() {
	r := gin.Default()
	// 中间件

	// 注册前台路由
	indexRouter := r.Group("")
	routers.ResistIndexRoutes(indexRouter)

	// 注册后端路由
	adminRouter := r.Group("/admin")
	routers.RegisterAdminRouter(adminRouter)
	r.Run("0.0.0.0:8800")
}
