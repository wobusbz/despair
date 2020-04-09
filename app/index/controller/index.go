package controller

import (
	"despair/Db"
	"despair/common/response"
	"github.com/gin-gonic/gin"
)

type IndexController struct {
}

func (index *IndexController) Index(ctx *gin.Context) {
	response.HttpResponse(ctx, gin.H{"messages": "hello world", "dbconnect": Db.DbConn.NewOrm().Ping()}, nil)
}
