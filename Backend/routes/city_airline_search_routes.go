package routes

import (
	"flight-companion-backend/controllers"
	"github.com/gin-gonic/gin"
)

// SetupRoutesForCityandAirlineSearch defines the city search route
func SetupRoutesForCityAndAirlineSearch(router *gin.RouterGroup) {
	router.GET("/search-city", controllers.SearchCityController)
	router.GET("/search-airline", controllers.SearchAirlineController)
}
