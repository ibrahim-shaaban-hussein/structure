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
        log.Fatal(err)
    }

    // Check the connection
    err = client.Ping(context.Background(), nil)
    if err != nil {
        log.Fatal(err)
    }

    // List databases
    databases, err := client.ListDatabaseNames(context.Background(), nil)
    if err != nil {
        log.Fatal(err)
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

