package routes

import (
	UserAuth "flightbuddy-backend/controllers/user/auth"
	"github.com/gin-gonic/gin"
)

// SetupUserRoutesForRegistrationAndLogin defines routes for user authentication
func SetupUserRoutesForRegistrationAndLogin(router *gin.RouterGroup) {
	router.POST("/signup", UserAuth.RegisterUser)
	router.POST("/login", UserAuth.LoginUser)
}
