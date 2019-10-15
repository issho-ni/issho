package mongo

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

// Client defines the interface for a service's MongoDB connection.
type Client interface {
	Collection(string, ...*options.CollectionOptions) *mongo.Collection
	Connect() context.CancelFunc
	Database() *mongo.Database
	DefineIndexes(...IndexSet)
}

type client struct {
	name string
	log  *log.Entry
	*mongo.Client
	indexSets []IndexSet
}

// NewClient creates a new MongoDB client for connecting to the specified database.
func NewClient(dbName string) Client {
	c := &client{name: dbName}

	c.log = log.WithFields(log.Fields{
		"mongodb.service": dbName,
		"span.kind":       "client",
		"system":          "mongodb",
	})

	uri := os.Getenv("MONGODB_URL")
	if uri == "" {
		c.log.Fatal("MongoDB URL must be set")
	}

	cc, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		c.log.WithField("err", err).Fatal("Failed to create MongoDB client")
	}

	c.Client = cc
	return c
}

// Collection returns the named collection on the MongoDB database.
func (c *client) Collection(name string, opts ...*options.CollectionOptions) *mongo.Collection {
	return c.Database().Collection(name, opts...)
}

// Connect establishes a connection to the MongoDB server.
func (c *client) Connect() context.CancelFunc {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	if err := c.Client.Connect(ctx); err != nil {
		c.log.WithField("err", err).Fatal("Failed to find MongoDB")
	}

	c.log.Debug("Connecting to MongoDB")
	pingCtx, cancelPing := context.WithTimeout(context.Background(), 30*time.Second)

	if err := c.Client.Ping(pingCtx, nil); err != nil {
		c.log.WithField("err", err).Fatal("Failed to connect to MongoDB")
	}

	cancelPing()
	c.log.Debug("Connected to MongoDB")

	c.createIndexes()

	return cancel
}

// Database returns the MongoDB database.
func (c *client) Database() *mongo.Database {
	return c.Client.Database(c.name)
}

// DefineIndexes specifies indexes to create on the database on connection.
func (c *client) DefineIndexes(indexSets ...IndexSet) {
	c.indexSets = append(c.indexSets, indexSets...)
}

func (c *client) createIndexes() {
	c.log.Debug("Creating indexes")
	createOptions := options.CreateIndexes().SetMaxTime(10 * time.Second)

	for _, indexSet := range c.indexSets {
		iv := c.Collection(indexSet.Collection).Indexes()
		results, err := iv.CreateMany(context.Background(), indexSet.Indexes, createOptions)
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
