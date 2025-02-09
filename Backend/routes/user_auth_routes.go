package routes

import (
	"flight-companion-backend/controllers"
	"github.com/gin-gonic/gin"
)

// SetupUserRoutesForRegistrationAndLogin defines routes for user authentication
func SetupUserRoutesForRegistrationAndLogin(router *gin.RouterGroup) {
	router.POST("/signup", controllers.SignUpUserController)
	router.POST("/login", controllers.LoginUserController)
}
