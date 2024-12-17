package main

import (
	"log"
	"presolai/configs"
	"presolai/internal/pkg/mysql"
	"presolai/internal/router"
)

func main() {

	configs.Init()
	EnvConfig := configs.EnvConfig

	log.Printf("MySQL Config: %+v\n", EnvConfig.Mysql)

	mysql.Connect(&EnvConfig.Mysql)

	r := router.InitRouter()

	// 启动服务器并监听请求
	log.Printf("👻 Server is now running on http://localhost%s\n", EnvConfig.Server.Port)
	if err := r.Run(EnvConfig.Server.Port); err != nil {
		log.Fatal("Unable to start server: ", err)
	}
}
