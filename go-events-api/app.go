package main

import (
	"example.com/go_events_api/pkgs/db"
	"example.com/go_events_api/pkgs/routes"
	"github.com/gin-gonic/gin"
)

func main() {

	db.InitDB()

	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":8080") //localhost:8080

}
