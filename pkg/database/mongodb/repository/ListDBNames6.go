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
    var databases []mongo.DatabaseInfo
    err = databasesCursor.All(ctx, &databases)
    if err != nil {
        log.Fatal("Error retrieving database names:", err)
    }

    // Print the list of databases
    log.Println("Databases:")
    for _, db := range databases {
        log.Println("-", db.Name)
    }

    // Close the connection when done
    err = client.Disconnect(context.Background())
    if err != nil {
        log.Fatal("Error disconnecting from MongoDB:", err)
    }
    log.Println("Connection closed")
}

