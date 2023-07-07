package middleware

import "github.com/gin-gonic/gin"

// AuthRolesMiddle 用于检测用户身份的
func AuthRolesMiddle() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 判断的逻辑

		ctx.Next()
	}
}
