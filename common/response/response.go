package response

import (
	"despair/common/errorMsg"
	"github.com/gin-gonic/gin"
	"net/http"
)

func HttpResponse(ctx *gin.Context, res, err interface{}) {
	baseErr, ok := err.(BaseError)
	if ok {
		ctx.JSON(http.StatusOK, gin.H{"code": baseErr.Code(), "messages": baseErr.Error()})
	}

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": errorMsg.ERROR_PARAM})
	}
	ctx.JSON(http.StatusOK, gin.H{"data": res})
}
