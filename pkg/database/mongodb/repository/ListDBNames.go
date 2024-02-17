package main

import (
	"context"
	"log"

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
	defer client.Disconnect(context.Background())

	// Check the connection
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal("Error pinging MongoDB:", err)
	}
	log.Println("Connected to MongoDB")

	// Get admin database
	adminDB := client.Database("admin")

	// List databases
	databasesCursor, err := adminDB.ListDatabases(context.Background(), nil)
	if err != nil {
		log.Fatal("Error listing databases:", err)
	}

	// Iterate over the databases
	var dbNames []string
	for databasesCursor.Next(context.Background()) {
		var dbInfo mongo.DatabaseInfo
		err := databasesCursor.Decode(&dbInfo)
		if err != nil {
			log.Fatal("Error decoding database info:", err)
		}
		dbNames = append(dbNames, dbInfo.Name)
	}

	// Print the list of databases
	log.Println("Databases:")
	for _, dbName := range dbNames {
		log.Println("-", dbName)
	}
}

