package citysearch

import (
	"context"
	"flightbuddy-backend/database"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

// CitySearchRequest represents the expected JSON body
type CitySearchRequest struct {
	City string `json:"city"`
}

// CityResult represents the structure of the city search results
type CityResult struct {
	City    string `bson:"city"`
	Country string `bson:"country"`
	Airport string `bson:"name"`
	IATA    string `bson:"iata_code"`
}

// SearchCity returns up to 10 city names matching user input
func SearchCity(c *gin.Context) {

	var request CitySearchRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request format. Expected JSON with 'city' key."})
		return
	}

	cityName := request.City
	log.Println(cityName)

	if cityName == "" {
		c.JSON(400, gin.H{"error": "City name is required"})
		return
	}

	collection := database.GetCollection("flightbuddy", "airports")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Case-insensitive partial match
	filter := bson.M{"city": bson.M{"$regex": "^" + cityName, "$options": "i"}}
	optionsToFilter := options.Find().SetLimit(10) // Limit results to 10

	cursor, err := collection.Find(ctx, filter, optionsToFilter)

	if err != nil {
		log.Println("Error searching for city:", err)
		c.JSON(500, gin.H{"error": "Failed to search city"})
		return
	}
	defer cursor.Close(ctx)

	var results []CityResult
	if err := cursor.All(ctx, &results); err != nil {
		log.Println("Error decoding search results:", err)
		c.JSON(500, gin.H{"error": "Failed to decode results"})
		return
	}

	c.JSON(200, results)
}
