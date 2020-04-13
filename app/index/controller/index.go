package controller

import (
	"despair/common/response"

	"github.com/gin-gonic/gin"
)

/*** 测试****/
type IndexRequest struct {
	Id int64 `json:"id"  binding:"required"`
}

type IndexController struct {
	//BaseController
}

func (index *IndexController) Index(ctx *gin.Context) {
	indexRequest := new(IndexRequest)
	if err := ctx.ShouldBindJSON(indexRequest); err != nil {

		response.HttpResponse(ctx, nil, err.Error())
	}

	//result, err := db.DbConn.NewOrm().QueryString("select now()")
	response.HttpResponse(ctx, indexRequest, nil)
}
