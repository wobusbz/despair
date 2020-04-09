package controller

import (
	"despair/Db"
	"github.com/gin-gonic/gin"
	"net/http"
)

type IndexController struct {
}

func (index IndexController) RegisterRoute(router *gin.RouterGroup) {
	router.GET("/", index.Index)
}

func (index *IndexController) Index(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"messages": "hello world", "dbconnect": Db.DbConn.NewOrm().Ping()})
}
