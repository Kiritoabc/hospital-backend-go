package router

import (
	"github.com/gin-gonic/gin"
	"hospital-backend-go/controller"
	"hospital-backend-go/middleware"
	"net/http"
)

// CollectRoute 路由控制器
func CollectRoute(r *gin.Engine) *gin.Engine {

	r.Use(middleware.CORSMiddleware())

	// 创建控制器实例
	Controller := &controller.Controller{}

	// 路由配置
	r.GET("/", func(context *gin.Context) {
		context.String(http.StatusOK, "Hello Gin")
	})

	// User的路由
	userRoutes := r.Group("/user")
	userRoutes.POST("/register", controller.Register)
	userRoutes.POST("/login", controller.Login)
	userRoutes.GET("/info", middleware.AuthMiddleware(), controller.GetUserInfo)

	//roomRoutes := r.Group("/room", middleware.AuthMiddleware())
	roomRoutes := r.Group("/room")
	roomRoutes.GET("/list", Controller.Room.RoomList)
	roomRoutes.POST("/delete", Controller.Room.DeleteById)
	roomRoutes.POST("/add", Controller.Room.AddRoom)
	roomRoutes.POST("/update", Controller.Room.Update)

	//doctorRoutes := r.Group("/doctor", middleware.AuthMiddleware())
	doctorRoutes := r.Group("/doctor")
	doctorRoutes.GET("/list", Controller.Doctor.DoctorList)
	doctorRoutes.POST("/delete", Controller.Doctor.DeleteById)
	doctorRoutes.POST("/add", Controller.Doctor.AddDoctor)
	doctorRoutes.POST("/update", Controller.Doctor.Update)
	return r
}
