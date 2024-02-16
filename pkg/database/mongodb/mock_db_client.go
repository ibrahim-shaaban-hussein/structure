package mongodb

import (
    "context"
    "fmt"
    "time"

    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "go.mongodb.org/mongo-driver/bson"
)

// MockDBClient is a mock implementation of a MongoDB client.
type MockDBClient struct {
    client *mongo.Client
}

// Connect simulates the connection to the MongoDB database.
func (c *MockDBClient) Connect() error {
    // Replace this with your actual MongoDB connection string
    connectionString := "mongodb://localhost:27017"

    // Set client options
    clientOptions := options.Client().ApplyURI(connectionString)

    // Connect to MongoDB
    client, err := mongo.Connect(context.Background(), clientOptions)
    if err != nil {
        return fmt.Errorf("error connecting to MongoDB: %v", err)
    }

    // Check the connection
    err = client.Ping(context.Background(), nil)
    if err != nil {
        return fmt.Errorf("error pinging MongoDB: %v", err)
    }

    c.client = client

    return nil
}

// InsertOne simulates inserting a document into the MongoDB collection.
func (c *MockDBClient) InsertOne() error {
    // Get the collection
    collection := c.client.Database("test").Collection("example")

    // Sample document to insert
    document := bson.M{"name": "John", "age": 30}

    // Insert document
    _, err := collection.InsertOne(context.Background(), document)
    if err != nil {
        return fmt.Errorf("error inserting document: %v", err)
    }

    return nil
}

// Close simulates closing the connection to the MongoDB database.
func (c *MockDBClient) Close() error {
    // Disconnect from MongoDB
    err := c.client.Disconnect(context.Background())
    if err != nil {
        return fmt.Errorf("error disconnecting from MongoDB: %v", err)
    }

    // Wait for the connection to close
    time.Sleep(1 * time.Second)

    return nil
}

