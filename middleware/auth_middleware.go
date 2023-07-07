package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
	"hospital-backend-go/database"
	"hospital-backend-go/model"
	"hospital-backend-go/service"
	"hospital-backend-go/util"
	"net/http"
	_ "net/http"
	"strings"
	_ "time"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 获取 authorization header
		tokenString := ctx.GetHeader("Authorization")

		fmt.Print("请求token", tokenString)

		//validate token formate
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "权限不足",
			})
			ctx.Abort()
			return
		}
		//tokenString = tokenString[7:] //截取字符
		claims, err := util.ParseTokenHs256(tokenString)

		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "权限不足",
			})
			ctx.Abort()
			return
		}
		//token通过验证, 获取claims中的phone
		phone := claims.Phone
		DB := database.GetDB()
		var doctor model.Doctor
		doctor, err = service.SelectDoctorByPhone(DB, phone)
		// 验证用户是否存在
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "权限不足",
			})
			ctx.Abort()
			return
		}
		//用户存在 将user信息写入上下文
		ctx.Set("doctorInfo", doctor)
		ctx.Next()
	}
}
