package main

import (
	"golearn/first-api/db"
	"golearn/first-api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()
  
  routes.RegisterRoutes(server)
	server.Run(":8080")
}
