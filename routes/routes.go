package routes

import (
	"eventBooking/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	// register a handler, routes
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)

	requireAuthentication := server.Group("/")
	requireAuthentication.Use(middlewares.Authenticate)
	requireAuthentication.POST("/events", createEvent) //protected'
	requireAuthentication.PUT("/events/:id", updateEvent)
	requireAuthentication.DELETE("/events/:id", deleteEvent)

	requireAuthentication.POST("/events/:id/register", registerForEvent)
	requireAuthentication.DELETE("/events/:id/register", cancelRegistration)

	server.POST("/signup", signUp)
	server.POST("/login", login)
	server.GET("/users", getUsers)

}
