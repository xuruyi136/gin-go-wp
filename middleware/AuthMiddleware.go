package middleware

import (
	"gin-go-wp/common"
	"gin-go-wp/model"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//获取authorization header
		tokenString := ctx.GetHeader("Authorization")
		//验证格式,不是以开头
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer") {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 400, "msg": "权限不足"})
			ctx.Abort()
			return
		}
		tokenString = tokenString[7:]
		token, claims, err := common.ParseToken(tokenString)
		if err != nil || !token.Valid {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 400, "msg": "权限不足"})
			ctx.Abort()
			return
		}
		//验证过后获取claim中的userId
		userId := claims.UserId
		DB := common.GetDB()
		var user model.User
		DB.First(&user, userId)
		//用户
		if user.ID == 0 {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 400, "msg": "权限不足"})
			ctx.Abort()
			return
		}
		//用户存在将user信息写入上下文
		ctx.Set("user", user)
		ctx.Next()

	}

}
