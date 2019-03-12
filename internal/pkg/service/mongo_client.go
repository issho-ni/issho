package service

import (
	"context"
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
	uri := os.Getenv("MONGODB_URL")

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
	log.Debug("Connected to MongoDB")

	ctx, cancelPing := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancelPing()

	err = c.Client.Ping(ctx, nil)
	if err != nil {
		log.Fatalf("Failed to ping MongoDB: %v", err)
	}
	log.Debug("Pinged MongoDB")

	return cancel
}

// Database returns the MongoDB database.
func (c *mongoClient) Database() *mongo.Database {
	return c.Client.Database(c.name)
}
