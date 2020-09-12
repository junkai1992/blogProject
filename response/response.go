package response

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-playground/validator/v10"
	"net/http"
)

func Response(ctx *gin.Context, httpStatus int, code int, data gin.H, msg interface{}) {
	ctx.JSON(httpStatus, gin.H{"code": code, "result": data, "msg": msg})
}

// 200 响应
func Success(ctx *gin.Context, data gin.H, msg string) {
	Response(ctx, http.StatusOK, 200, data, msg)
}

// 400 响应
func Fail(ctx *gin.Context, msg string, data gin.H) {
	Response(ctx, http.StatusBadRequest, 400, data, msg)
}
