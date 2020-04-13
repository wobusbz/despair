package middleware

import (
	"despair/common/jwt"
	"despair/common/response"
	"despair/global"
	"github.com/gin-gonic/gin"
)

// **************验证token路由中间件***********

func AuthMiddleware(auth *jwt.AuthToken) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("Authorization")

		if token == "" {
			response.HttpResponse(ctx, nil, response.NewBaseError(global.ERROR_TOKEN_CODE, global.ERROR_TOKEN_MSG))
			ctx.Abort()
			return
		}
		authToken, err := auth.DecodeToken(token)
		if err != nil {
			response.HttpResponse(ctx, nil, response.NewBaseError(global.ERROR_DECODE_CODE, err.Error()))
			ctx.Abort()
			return
		}
		ctx.Set("userAuth", authToken)
	}
}
