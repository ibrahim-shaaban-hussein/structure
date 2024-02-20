package repository

import (
    "context"
    "log"

    "go.mongodb.org/mongo-driver/mongo"
)

func ListDatabaseNames(client *mongo.Client) {
    // Assuming adminDB is correctly initialized
    // List database names
    dbNames, err := client.ListDatabaseNames(context.Background(), nil)
    if err != nil {
        log.Fatal(err)
    }

    // Print the list of database names
    log.Println("Database Names:")
    for _, dbName := range dbNames {
        log.Println("-", dbName)
    }
}

