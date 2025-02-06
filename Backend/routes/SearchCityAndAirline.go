package routes

import (
	"flightbuddy-backend/controllers/arlinesearch"
	"flightbuddy-backend/controllers/citysearch"
	"github.com/gin-gonic/gin"
)

// SetupRoutesForCityandAirlineSearch defines the city search route
func SetupRoutesForCityAndAirlineSearch(router *gin.RouterGroup) {
	router.GET("/search-city", citysearch.SearchCity)
	router.GET("/search-airline", arlinesearch.SearchAirline)
}
