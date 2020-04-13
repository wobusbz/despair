package main

import (
	"despair/app/index/controller"
	"despair/app/middleware"
	"despair/app/routers"
	"despair/common/logger"
	"despair/db"
	"github.com/gin-gonic/gin"
)

func init() {
	logger.InitLogger()
	db.InitMysql("mysql", "root:@tcp(127.0.0.1:3306)/wuhuarou?charset=utf8&parseTime=True&loc=Local")
	logger.Debug("mysql connect succ")
}

func main() {

	r := gin.New()
	// 中间件
	// 初始化基础模型，激活auth认证
	// 全局初始化一个auth
	baseController := controller.NewBaseController()
	r.Use(middleware.AuthMiddleware(baseController.Auth))

	// 注册前台路由
	indexRouter := r.Group("")
	routers.ResistIndexRoutes(indexRouter)
	logger.Debug("abc%s", "abc")

	// 注册后端路由
	adminRouter := r.Group("/admin")
	routers.RegisterAdminRouter(adminRouter)
	r.Run("0.0.0.0:8800")
}
