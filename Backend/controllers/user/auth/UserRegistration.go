package auth

import (
	"context"
	"flightbuddy-backend/database"
	"flightbuddy-backend/structs"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"net/mail"
	"time"
)

func RegisterUser(c *gin.Context) {

	var user structs.UserSignUpDetails

	// Bind and validate input
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request. Ensure email is valid and password is at least 6 characters."})
		return
	}

	if !ValidateEmail(user.Email) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email format"})
		return
	}

	collection := database.GetCollection("flightbuddy", "users")
	if collection == nil {
		log.Println("Error: Failed to get MongoDB collection")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	// Create context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// **Check if email already exists**
	var existingUser structs.UserSignUpDetails

	err := collection.FindOne(ctx, bson.M{"email": user.Email}).Decode(&existingUser)

	if err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Email already registered"})
		return
	}

	// **Hash the password before storing**
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	// Insert user into MongoDB
	_, err = collection.InsertOne(ctx, bson.M{
		"email":     user.Email,
		"password":  string(hashedPassword), // Store hashed password
		"gender":    user.Gender,
		"languages": user.Languages,
		"createdAt": time.Now(),
	})

	if err != nil {
		log.Println("Error inserting user:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Signup successfull"})
}

// ValidateEmail checks if the email format is valid
func ValidateEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}
