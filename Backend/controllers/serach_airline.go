package controllers

import (
	"flight-companion-backend/service"
	"flight-companion-backend/structs"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SearchAirlineController(c *gin.Context) {

	var request structs.AirlineSearchRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format. Expected JSON with 'airline' key."})
		return
	}

	results, err := service.SearchAirline(request.Airline)
	if err != nil {
		if err.Error() == "airline name is required" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Airline name cannot be empty"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch airline results"})
		}
		return
	}

	// If no results found, return 404
	if len(results) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "No matching airlines found"})
		return
	}

	c.JSON(http.StatusOK, results)
}
