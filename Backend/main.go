package main

import (
	"flightbuddy-backend/database"
	"flightbuddy-backend/routes"
	"github.com/gin-gonic/gin"
)

func main() {

	database.ConnectToDB()
	// Create a new Gin router
	r := gin.Default()

	// Define a simple GET route
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, Gin!",
		})
	})

	routes.SetupAllRoutes(r)

	//routes.SetupUserRoutesForRegistrationAndLogin(r)
	//routes.SetupRoutesForCitySearch(r)
	//routes.GetUsersOnFlightRoute(r)
	//routes.SetUpAddUserToFlightRoute(r)
	// Start the server on port 8080
	r.Run(":8080")
}
