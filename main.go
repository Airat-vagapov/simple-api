package main

import (
	"example.com/simple-api/db"
	"example.com/simple-api/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()

	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":8080") // localhost:8080

}
