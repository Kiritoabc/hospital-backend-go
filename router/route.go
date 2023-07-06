package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// CollectRoute 路由控制器
func CollectRoute(r *gin.Engine) *gin.Engine {

	// 路由配置
	r.GET("/", func(context *gin.Context) {
		context.String(http.StatusOK, "Hello Gin")
	})

	return r
}
