package repository

import (
    "context"
    "log"

    "go.mongodb.org/mongo-driver/mongo"
)

func ListDatabases(client *mongo.Client) {
    // Assuming client is already connected
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
}

