package main

import (
	"github.com/gin-gonic/gin"
	"github.com/intellect-sam/event/db"
	"github.com/intellect-sam/event/routes"
)

func main() {
	db.InitDB()

	server := gin.Default()
	routes.RegisteerRoutes(server)

	server.Run(":8080")
}
