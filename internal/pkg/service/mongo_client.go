package service

import (
	"context"
	"fmt"
	"os"
	"time"

	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoClient defines the interface for a service's MongoDB connection.
type MongoClient interface {
	Connect() context.CancelFunc
	Database() *mongo.Database
}

type mongoClient struct {
	name string
	*mongo.Client
}

// NewMongoClient creates a new MongoDB client for connecting to the specified database.
func NewMongoClient(dbName string) MongoClient {
	uri := fmt.Sprintf("mongodb://%s:%s@issho-mongodb:27017", os.Getenv("MONGODB_USER"), os.Getenv("MONGODB_PASS"))

	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatalf("Failed to create MongoDB client: %v", err)
	}

	return &mongoClient{dbName, client}
}

// Connect establishes a connection to the MongoDB server.
func (c *mongoClient) Connect() context.CancelFunc {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	err := c.Client.Connect(ctx)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	return cancel
}

// Database returns the MongoDB database.
func (c *mongoClient) Database() *mongo.Database {
	return c.Client.Database(c.name)
}
