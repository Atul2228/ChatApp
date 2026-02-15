package config

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func ConnectDB(databaseUrl string) *mongo.Client {
	// Replace with your Atlas URI
	uri := databaseUrl

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}

	// Ping the database to verify connection
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("Could not connect to MongoDB Atlas:", err)
	}

	log.Println("Connected to MongoDB Atlas!")
	return client
}

// Helper to get a collection
func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	return client.Database("User").Collection(collectionName)
}
