package routers

import (
	"despair/app/index/controller"
	"github.com/gin-gonic/gin"
)

func ResistIndexRoutes(r *gin.RouterGroup) {
	new(controller.IndexController).RegisterRoute(r)
}
