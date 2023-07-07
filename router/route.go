package router

import (
	"github.com/gin-gonic/gin"
	"hospital-backend-go/controller"
	"net/http"
)

// CollectRoute 路由控制器
func CollectRoute(r *gin.Engine) *gin.Engine {

	// 路由配置
	r.GET("/", func(context *gin.Context) {
		context.String(http.StatusOK, "Hello Gin")
	})

	// User的路由
	userRoutes := r.Group("/user")
	userRoutes.POST("/register", controller.Register)

	return r
}
