package controllers

import (
	"flightbuddy-backend/service"
	"flightbuddy-backend/structs"
	"github.com/gin-gonic/gin"
	"net/http"
)

// AddUserToFlightController handles adding a user to a flight.
func AddUserToFlightController(c *gin.Context) {

	// Extract user email from JWT Token
	userEmail, exists := c.Get("currentUserEmail")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	var request structs.FlightAddUserRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format. Required fields: source_iata, destination_iata, airline, date, departure_time"})
		return
	}

	message, err := service.AddUserToFlight(userEmail.(string), request)
	if err != nil {
		if err.Error() == "user already exists" {
			c.JSON(http.StatusConflict, gin.H{"error": "You are already added to this flight"})
		} else if err.Error() == "failed to add user" {
			c.JSON(http.StatusCreated, gin.H{"message": "Failed to add you to the flight"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Server error while adding, please try again later!"})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": message})
}
