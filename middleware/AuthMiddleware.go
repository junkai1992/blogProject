package middleware

import (
	"blogBack/common"
	"blogBack/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 获取authorization
		tokenString := ctx.GetHeader("Authorization")
		fmt.Println(tokenString)
		// 校验:如果没有token或token不是Bearer开头，返回权限不足。
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足"})
			ctx.Abort()
			return
		}
		// 去除前面 Bearer
		tokenString = tokenString[7:]
		// 解析token
		token, claims, err := common.ParseToken(tokenString)
		// 如果有错误，或token无法通过验证，返回权限不足
		if err != nil || !token.Valid {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足"})
			ctx.Abort()
			return
		}
		// 验证通过后获取claim,中userId
		userId := claims.UserId
		// 验证用户是否在数据库存在
		DB := common.GetDB()
		var user model.User
		DB.First(&user, userId)
		// 查询不到会返回0
		if user.ID == 0 {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足"})
			ctx.Abort()
			return
		}
		// 用户存在将user信息写入上下文
		ctx.Set("user", user)
		ctx.Next()
	}
}
