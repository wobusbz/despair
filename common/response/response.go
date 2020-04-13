package response

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/go-playground/v10"
	"net/http"
)

func HttpResponse(ctx *gin.Context, res, err interface{}) {
	baseErr, ok := err.(BaseError)
	if ok {
		ctx.JSON(http.StatusOK, gin.H{"code": baseErr.Code(), "msg": baseErr.Error()})
		ctx.Abort()
		return
	}

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": err})
		ctx.Abort()
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": res})
}

func ParmError(ctx *gin.Context, err interface{}) {
	validErr, ok := err.(validator.ValidationErrors)
	if ok {
		fmt.Println(validErr)
	}

	fmt.Println(validErr)
}
