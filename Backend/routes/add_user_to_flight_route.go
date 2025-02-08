package routes

import (
	"flightbuddy-backend/controllers"
	"github.com/gin-gonic/gin"
)

func SetUpAddUserToFlightRoute(router *gin.RouterGroup) {

	router.POST("/add-user", controllers.AddUserToFlightController)

}
