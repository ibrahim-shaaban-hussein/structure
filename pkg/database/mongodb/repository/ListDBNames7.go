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
		log.Fatal("Error connecting to MongoDB:", err)
	}

	// Check the connection
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("Error pinging MongoDB:", err)
	}

	// List databases
	databasesCursor, err := client.ListDatabases(ctx, nil)
	if err != nil {
		log.Fatal("Error listing databases:", err)
	}

	// Iterate over the cursor and print the list of databases
	var databases []string
	for databasesCursor.Next(ctx) {
		var db mongo.DatabaseInfo
		if err := databasesCursor.Decode(&db); err != nil {
			log.Fatal("Error decoding database:", err)
		}
		databases = append(databases, db.Name)
	}
	if err := databasesCursor.Err(); err != nil {
		log.Fatal("Error iterating over database cursor:", err)
	}

	// Print the list of databases
	log.Println("Databases:")
	for _, db := range databases {
		log.Println("-", db)
	}

	// Close the connection when done
	err = client.Disconnect(context.Background())
	if err != nil {
		log.Fatal("Error disconnecting from MongoDB:", err)
	}
	log.Println("Connection closed")
}

