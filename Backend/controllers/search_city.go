package controllers

import (
	"flightbuddy-backend/service"
	"flightbuddy-backend/structs"
	"github.com/gin-gonic/gin"
	"net/http"
)

// SearchCity returns up to 10 city names matching user input
func SearchCityController(c *gin.Context) {

	var request structs.CitySearchRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request format. Expected JSON with 'city' key."})
		return
	}

	results, err := service.SearchCity(request.City)

	if err != nil {
		if err.Error() == "city name is required" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "City name cannot be empty"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch city results"})
		}
		return
	}

	// If no results found, return 404
	if len(results) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "No matching city is found"})
		return
	}

	c.JSON(http.StatusOK, results)
}
