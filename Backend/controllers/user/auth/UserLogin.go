package auth

import (
	"context"
	"flightbuddy-backend/database"
	"flightbuddy-backend/jwt"
	"flightbuddy-backend/structs"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"time"
)

// LoginUser handles user login
func LoginUser(c *gin.Context) {

	var loginData structs.UserLoginDetails

	if err := c.ShouldBindJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request. Ensure email and password are provided."})
		return
	}

	collection := database.GetCollection("flightbuddy", "users")
	if collection == nil {
		log.Println("Error: Failed to get MongoDB collection")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	projection := bson.M{"email": 1, "password": 1, "_id": 0} // Only fetch email & password

	var user structs.UserLoginDetails
	err := collection.FindOne(ctx, bson.M{"email": loginData.Email}, options.FindOne().SetProjection(projection)).Decode(&user)

	//err := collection.FindOne(ctx, bson.M{"email": loginData.Email}).Decode(&userFound)

	if err == mongo.ErrNoDocuments {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "You are not a registered user!"})
		return
	}

	// Compare hashed password from DB with the input password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginData.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Wrong password"})
		return
	}

	token, err := jwt.GenerateJWT(user.Email)
	if err != nil {
		log.Println("Error generating JWT token:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"token":   token,
	})
}
