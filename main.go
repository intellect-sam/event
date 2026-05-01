// @title           Event API
// @version         1.0
// @description     API for managing events and users
// @host            localhost:8080
// @BasePath        /

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

package main

import (
	"github.com/gin-gonic/gin"
	"github.com/intellect-sam/event/db"
	_ "github.com/intellect-sam/event/docs"
	"github.com/intellect-sam/event/routes"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	db.InitDB()

	server := gin.Default()
	routes.RegisteerRoutes(server)

	// Swagger endpoint
	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	server.Run(":8080")
}
