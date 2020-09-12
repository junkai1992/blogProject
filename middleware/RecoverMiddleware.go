package middleware

import (
	"blogBack/response"
	"fmt"
	"github.com/gin-gonic/gin"
)

// 全局异常捕获
func RecoveryMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				response.Fail(ctx, fmt.Sprint(err), nil)
			}
		}()
		ctx.Next()
	}
}
