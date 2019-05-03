package service

import (
	"context"
	"os"
	"time"

	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// IndexSet specifies the indexes to create for the given collection.
type IndexSet struct {
	Collection string
	Indexes    []mongo.IndexModel
}

// NewIndexSet creates a new IndexSet.
func NewIndexSet(collection string, indexes ...mongo.IndexModel) IndexSet {
	return IndexSet{Collection: collection, Indexes: indexes}
}

// MongoClient defines the interface for a service's MongoDB connection.
type MongoClient interface {
	Collection(string, ...*options.CollectionOptions) *mongo.Collection
	Connect() context.CancelFunc
	CreateIndexes(...IndexSet)
	Database() *mongo.Database
}

type mongoClient struct {
	name string
	log  *log.Entry
	*mongo.Client
}

// NewMongoClient creates a new MongoDB client for connecting to the specified database.
func NewMongoClient(dbName string) MongoClient {
	uri := os.Getenv("MONGODB_URL")

	entry := log.WithFields(log.Fields{
		"mongodb.service": dbName,
		"span.kind":       "client",
		"system":          "mongodb",
	})

	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		entry.WithField("err", err).Fatal("Failed to create MongoDB client")
	}

	return &mongoClient{dbName, entry, client}
}

// Connect establishes a connection to the MongoDB server.
func (c *mongoClient) Connect() context.CancelFunc {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	err := c.Client.Connect(ctx)
	if err != nil {
		c.log.WithField("err", err).Fatal("Failed to find MongoDB")
	}

	c.log.Debug("Connecting to MongoDB")
	pingCtx, cancelPing := context.WithTimeout(context.Background(), 30*time.Second)

	err = c.Client.Ping(pingCtx, nil)
	if err != nil {
		c.log.WithField("err", err).Fatal("Failed to connect to MongoDB")
	}

	c.log.Debug("Connected to MongoDB")
	cancelPing()

	return cancel
}

// CreateIndexes creates the specified indexes on the client's database.
func (c *mongoClient) CreateIndexes(indexSets ...IndexSet) {
	c.log.Debug("Creating indexes")
	createOptions := options.CreateIndexes().SetMaxTime(10 * time.Second)

	for _, indexSet := range indexSets {
		coll := c.Collection(indexSet.Collection).Indexes()

		results, err := coll.CreateMany(context.Background(), indexSet.Indexes, createOptions)
		if err != nil {
			c.log.WithField("err", err).Fatal("Could not create indexes")
		}

		for _, result := range results {
			c.log.WithFields(log.Fields{
				"mongodb.collection": indexSet.Collection,
				"mongodb.index":      result,
			}).Debug("Created index")
		}
	}
}

// Database returns the MongoDB database.
func (c *mongoClient) Database() *mongo.Database {
	return c.Client.Database(c.name)
}

// Collection returns the named collection on the MongoDB database.
func (c *mongoClient) Collection(name string, opts ...*options.CollectionOptions) *mongo.Collection {
	return c.Database().Collection(name, opts...)
}
