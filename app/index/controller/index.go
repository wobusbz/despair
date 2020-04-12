package controller

import (
	"despair/common/response"
	"despair/db"
	"github.com/gin-gonic/gin"
)

type IndexController struct {
}

func (index *IndexController) Index(ctx *gin.Context) {
	result, err := db.DbConn.NewOrm().QueryString("select now()")
	if err != nil {
		response.HttpResponse(ctx, nil, err)
		return
	}
	response.HttpResponse(ctx, result, nil)
}
