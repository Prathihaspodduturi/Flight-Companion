package routes

import (
	"flight-companion-backend/controllers"
	"github.com/gin-gonic/gin"
)

func SetUpAddUserToFlightRoute(router *gin.RouterGroup) {

	router.POST("/add-user", controllers.AddUserToFlightController)

}
