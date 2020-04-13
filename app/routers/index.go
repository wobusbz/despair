package routers

import (
	"despair/app/index/controller"
	"despair/app/middleware"
	"github.com/gin-gonic/gin"
)

func ResistIndexRoutes(r *gin.RouterGroup) {
	// 初始化基础模型，激活auth认证
	// 全局初始化一个auth
	baseController := controller.NewBaseController()
	middleware.AuthMiddleware(baseController.Auth)
	//*************index**********************

	r.GET("/index", new(controller.IndexController).Index)

}
