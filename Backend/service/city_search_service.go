package service

import (
	"context"
	"errors"
	"flightbuddy-backend/database"
	"flightbuddy-backend/structs"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

// SearchCity handles city search logic and returns up to 10 matching city names.
func SearchCity(cityName string) ([]structs.CityResult, error) {
	if cityName == "" {
		return nil, errors.New("city name is required")
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
		return nil, errors.New("database error")
	}
	defer cursor.Close(ctx)

	var results []structs.CityResult

	if err := cursor.All(ctx, &results); err != nil {
		log.Println("Error decoding search results:", err)
		return nil, errors.New("failed to decode results")
	}

	return results, nil
}
