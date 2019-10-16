package main

import (
	"go-template-rest-api/config"
	"go-template-rest-api/src/database"
	"go-template-rest-api/src/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {

	// config
	conf := new(config.ConfigReader)
	conf.Read()

	// DB
	database.InitMysql()

	// routes
	defaultRoute := new(routes.DefaultRoute)
	router := gin.Default()
	defaultRoute.InitRoute(router)

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	router.Use(cors.New(corsConfig))

	router.Run(viper.GetString("server.host") + ":" + viper.GetString("server.port"))
}
