package main

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Connect to MongoDB
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	// List databases
	ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	cursor, err := client.ListDatabases(ctx, nil)
	if err != nil {
		log.Fatal("Error listing databases:", err)
	}

	// Iterate over the cursor and print the database names
	log.Println("Databases:")
	for cursor.Next(context.Background()) {
		var result struct {
			Name string `bson:"name"`
		}
		err := cursor.Decode(&result)
		if err != nil {
			log.Fatal("Error decoding database name:", err)
		}
		log.Println("-", result.Name)
	}
	if err := cursor.Err(); err != nil {
		log.Fatal("Cursor error:", err)
	}

	// Close the cursor
	if err := cursor.Close(context.Background()); err != nil {
		log.Fatal("Error closing cursor:", err)
	}

	// Close the connection when done
	err = client.Disconnect(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connection closed")
}

