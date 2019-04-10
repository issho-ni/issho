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
	Collection(string, ...*options.CollectionOptions) *mongo.Collection
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
		log.Fatalf("Failed to find MongoDB: %v", err)
	}

	log.Debug("Connecting to MongoDB")
	pingCtx, cancelPing := context.WithTimeout(context.Background(), 30*time.Second)

	err = c.Client.Ping(pingCtx, nil)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	log.Debug("Connected to MongoDB")
	cancelPing()

	return cancel
}

// Database returns the MongoDB database.
func (c *mongoClient) Database() *mongo.Database {
	return c.Client.Database(c.name)
}

// Collection returns the named collection on the MongoDB database.
func (c *mongoClient) Collection(name string, opts ...*options.CollectionOptions) *mongo.Collection {
	return c.Database().Collection(name, opts...)
}
