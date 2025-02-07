package controllers

import (
	"flightbuddy-backend/service"
	"flightbuddy-backend/structs"
	"github.com/gin-gonic/gin"
	"net/http"
)

// LoginUser handles user login
func LoginUserController(c *gin.Context) {

	var loginData structs.UserLoginDetails

	if err := c.ShouldBindJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request. Ensure email and password are provided."})
		return
	}

	token, err := service.LoginUser(loginData)
	if err != nil {
		switch err.Error() {
		case "user not found":
			c.JSON(http.StatusNotFound, gin.H{"error": "User not registered"})
		case "wrong password":
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Incorrect password"})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"token":   token,
	})
}
