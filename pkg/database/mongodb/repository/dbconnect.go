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

    // Database and collection names
    dbName := "mydatabase"
    collectionName := "mycollection"

    // Create a new database
    db := client.Database(dbName)

    // Create a new collection
    collection := db.Collection(collectionName)

    // You can now perform operations on the collection, such as inserting documents
    // For example:
    // _, err = collection.InsertOne(context.Background(), bson.M{"name": "John Doe", "age": 30})
    // if err != nil {
    //     log.Fatal(err)
    // }

    // Close the connection when done
    err = client.Disconnect(context.Background())
    if err != nil {
        log.Fatal(err)
    }
    log.Println("Connection closed")
}

