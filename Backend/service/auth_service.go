package service

import (
	"context"
	"errors"
	"flightbuddy-backend/database"
	"flightbuddy-backend/jwt"
	"flightbuddy-backend/structs"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/mail"
	"strings"
	"time"
)

// RegisterUser handles user registration logic.
func RegisterUser(user structs.UserSignUpDetails) error {

	user.Email = strings.ToLower(user.Email)

	if !validateEmail(user.Email) {
		return errors.New("invalid email format")
	}

	collection := database.GetCollection("flightbuddy", "users")
	if collection == nil {
		log.Println("Error: Failed to get MongoDB collection")
		return errors.New("internal server error")
	}

	// Create context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// **Check if email already exists**
	var existingUser structs.UserSignUpDetails
	err := collection.FindOne(ctx, bson.M{"email": user.Email}).Decode(&existingUser)

	if err == nil {
		return errors.New("email already registered")
	}

	// **Hash the password before storing**
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("Error hashing password:", err)
		return errors.New("failed to hash password")
	}

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
		return errors.New("failed to register user")
	}

	return nil
}

// LoginUser handles the login logic.
func LoginUser(loginData structs.UserLoginDetails) (string, error) {

	collection := database.GetCollection("flightbuddy", "users")
	if collection == nil {
		log.Println("Error: Failed to get MongoDB collection")
		return "", errors.New("internal server error")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Convert email to lowercase before querying
	email := strings.ToLower(loginData.Email)

	projection := bson.M{"email": 1, "password": 1, "_id": 0} // Only fetch email & password

	var user structs.UserLoginDetails
	err := collection.FindOne(ctx, bson.M{"email": email}, options.FindOne().SetProjection(projection)).Decode(&user)

	if err == mongo.ErrNoDocuments {
		return "", errors.New("user not found")
	}

	// Compare hashed password from DB with the input password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginData.Password))
	if err != nil {
		return "", errors.New("wrong password")
	}

	// Generate JWT token
	token, err := jwt.GenerateJWT(user.Email)
	if err != nil {
		log.Println("Error generating JWT token:", err)
		return "", errors.New("internal server error")
	}

	return token, nil
}

// validateEmail checks if the email format is valid.
func validateEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}
