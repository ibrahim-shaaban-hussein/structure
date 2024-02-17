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
	databases, err := client.ListDatabaseNames(ctx, nil)
	if err != nil {
		log.Fatal("Error listing databases:", err) // Print the error
	}

	// Check if databases slice is nil
	if databases == nil {
		log.Println("No databases found")
		return
	}

	// Print the list of databases
	log.Println("Databases:")
	for _, db := range databases {
		log.Println("-", db)
	}

	// Close the connection when done
	err = client.Disconnect(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connection closed")
}

