package details

import "go.mongodb.org/mongo-driver/mongo"

// MongoDetail implement detailer interface to work with mongoDB
type MongoDetail struct {
	collection *mongo.Collection
}

// New create new instance of MongoDetail with mongo collection
func New(c *mongo.Collection) *MongoDetail {
	return &MongoDetail{collection: c}
}
