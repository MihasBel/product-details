package mgdetailer

import "go.mongodb.org/mongo-driver/mongo"

// MongoDetailer implement detailer interface to work with mongoDB
type MongoDetailer struct {
	collection *mongo.Collection
}

// New create new instance of MongoDetailer with mongo collection
func New(c *mongo.Collection) *MongoDetailer {
	return &MongoDetailer{collection: c}
}
