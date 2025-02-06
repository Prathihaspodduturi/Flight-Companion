package arlinesearch

import (
	"context"
	"flightbuddy-backend/database"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

// AirlineSearchRequest represents the expected JSON body
type AirlineSearchRequest struct {
	Airline string `json:"airline"`
}

func SearchAirline(c *gin.Context) {

	var request AirlineSearchRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request format. Expected JSON with 'city' key."})
		return
	}

	airlineName := request.Airline
	log.Println("searching for airline : ", airlineName)

	if airlineName == "" {
		c.JSON(400, gin.H{"error": "Airline name is required"})
		return
	}

	collection := database.GetCollection("flightbuddy", "airlines")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Case-insensitive partial match
	filter := bson.M{"airline": bson.M{"$regex": "^" + airlineName, "$options": "i"}}
	options := options.Find().SetLimit(10) // Limit results to 10

	cursor, err := collection.Find(ctx, filter, options)

	if err != nil {
		log.Println("Error searching for airline:", err)
		c.JSON(500, gin.H{"error": "Failed to search airline"})
		return
	}
	defer cursor.Close(ctx)

	var results []AirlineSearchRequest
	if err := cursor.All(ctx, &results); err != nil {
		log.Println("Error decoding search results:", err)
		c.JSON(500, gin.H{"error": "Failed to decode results"})
		return
	}

	c.JSON(200, results)
}
