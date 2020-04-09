package routers

import (
	"despair/app/index/controller"
	"github.com/gin-gonic/gin"
)

func ResistIndexRoutes(r *gin.RouterGroup) {
	//*************index**********************
	r.GET("/", new(controller.IndexController).Index)
}
