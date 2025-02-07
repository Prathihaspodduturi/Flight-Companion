package database

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

var Client *mongo.Client

// ConnectToDB connects to MongoDB Atlas and sets the global Client variable
func ConnectToDB() {
	var err error
	// Replace <username>, <password>, <cluster-url>, and <dbname> with your actual details
	uri := "mongodb+srv://Cluster65042:FlightBuddy%403103@cluster65042.poqrq.mongodb.net/?retryWrites=true&w=majority&appName=Cluster65042"

	// Connect to MongoDB
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	Client, err = mongo.Connect(ctx, options.Client().ApplyURI(uri))

	if err != nil {
		log.Fatal("Error creating MongoDB client:", err)
	}

	if err = Client.Ping(ctx, nil); err != nil {
		log.Fatal("Failed to ping MongoDB:", err)
	}

	log.Println("Connected to MongoDB!")
}

// GetDatabase returns a MongoDB database
func GetDatabase(dbName string) *mongo.Database {

	if Client == nil {
		log.Println("MongoDB client is not connected")
		return nil
	}
	return Client.Database(dbName)

}

// GetCollection returns a collection from the specified database
func GetCollection(dbName, collectionName string) *mongo.Collection {

	if collectionName == "" {
		log.Println("collection name cannot be empty")
		return nil
	}

	db := GetDatabase(dbName)
	if db == nil {
		log.Println("failed to get database")
		return nil
	}

	return db.Collection(collectionName)
}
