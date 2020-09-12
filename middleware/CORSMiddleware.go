package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 设置允许访问域名
		//ctx.Writer.Header().Set("Access-Control-Allow-Origin","http://localhost:9216")
		ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		// 设置缓存时间
		ctx.Writer.Header().Set("Access-Control-Max-Age", "86400")
		// 允许请求方法
		ctx.Writer.Header().Set("Access-Control-Allow-Methods", "*")
		// 允许请求头信息
		ctx.Writer.Header().Set("Access-Control-Allow-Headers", "*")
		// 允许客户端携带验证头信息 如:Cookies
		ctx.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		// 如果请求是Option 返回200
		if ctx.Request.Method == http.MethodOptions {
			ctx.AbortWithStatus(200)
		} else {
			// 否则向下继续判定
			ctx.Next()
		}
	}
}
