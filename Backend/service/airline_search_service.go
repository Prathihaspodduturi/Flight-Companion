package service

import (
	"context"
	"errors"
	"flight-companion-backend/database"
	"flight-companion-backend/structs"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

// SearchAirline handles airline search logic and returns up to 10 matching airlines.
func SearchAirline(airlineName string) ([]structs.AirlineResult, error) {
	if airlineName == "" {
		return nil, errors.New("airline name is required")
	}

	collection := database.GetCollection("flightbuddy", "airlines")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Case-insensitive partial match
	filter := bson.M{"airline": bson.M{"$regex": "^" + airlineName, "$options": "i"}}
	optionsToFilter := options.Find().SetLimit(10) // Limit results to 10

	cursor, err := collection.Find(ctx, filter, optionsToFilter)
	if err != nil {
		log.Println("Error searching for airline:", err)
		return nil, errors.New("failed to search airline")
	}
	defer cursor.Close(ctx)

	var results []structs.AirlineResult
	if err := cursor.All(ctx, &results); err != nil {
		log.Println("Error decoding search results:", err)
		return nil, errors.New("failed to decode results")
	}

	return results, nil
}
