package main

import (
	"eventBooking/db"
	"eventBooking/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	// configures http serve
	server := gin.Default()
	routes.RegisterRoutes(server)

	server.Run(":8080") //localhost 8080

}
