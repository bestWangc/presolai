package main

import (
	ginSwagger "github.com/swaggo/gin-swagger"
	ginSwaggerFiles "github.com/swaggo/gin-swagger/swaggerFiles"
	"log"
	"presolai/configs"
	_ "presolai/docs"
	"presolai/internal/pkg/mysql"
	"presolai/internal/router"
)

func main() {

	configs.Init()
	EnvConfig := configs.EnvConfig

	log.Printf("MySQL Config: %+v\n", EnvConfig.Mysql)

	mysql.Connect(&EnvConfig.Mysql)

	r := router.InitRouter()

	// 配置 Swagger 文档路由
	r.GET("/swagger/*any", ginSwagger.WrapHandler(ginSwaggerFiles.Handler))

	// 启动服务器并监听请求
	log.Printf("👻 Server is now running on http://localhost%s\n", EnvConfig.Server.Port)
	if err := r.Run(EnvConfig.Server.Port); err != nil {
		log.Fatal("Unable to start server: ", err)
	}
}
