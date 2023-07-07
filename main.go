package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"hospital-backend-go/database"
	"hospital-backend-go/router"
	"hospital-backend-go/util"
	"log"
	"os"
)

func main() {
	// 加载配置文件
	InitConfig()
	// 初始化雪花算法
	if err := util.Init(viper.GetString("start_time"), viper.GetInt64("machine_id")); err != nil {
		fmt.Printf("init snowflake failed, err:%v\n", err)
		return
	}
	// 初始化DB
	err := database.InitDB()
	if err != nil {
		log.Fatalf("连接数据库失败: %v", err)
	}
	// 初始化路由
	r := gin.Default()
	r = router.CollectRoute(r)
	port := viper.GetString("server.port")
	if port != "" {
		panic(r.Run(":" + port))
	}
	// listen and serve on 0.0.0.0:8080
	panic(r.Run())
}

func InitConfig() {
	worDir, _ := os.Getwd()
	viper.SetConfigName("application")
	viper.SetConfigType("yml")
	viper.AddConfigPath(worDir + "/config")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("配置文件读取出错: %v", err)
	}
}
