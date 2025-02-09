package controllers

import (
	"flight-companion-backend/service"
	"flight-companion-backend/structs"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SignUpUserController(c *gin.Context) {

	var user structs.UserSignUpDetails

	// Bind and validate input
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request. Ensure email is valid and password is at least 6 characters."})
		return
	}

	err := service.RegisterUser(user)
	if err != nil {
		if err.Error() == "email already registered" {
			c.JSON(http.StatusConflict, gin.H{"error": "Email is already registered"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user"})
		}
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Signup successful"})
}
