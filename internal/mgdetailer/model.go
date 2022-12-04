package mgdetailer

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// details ID     primitive.ObjectID `json:"id" bson:"_id"`
// map[[2]string]string
// details model to represent details in database
type details struct {
	ID          primitive.ObjectID `json:"id,omitempty" bson:"_id"`
	ProductName string             `json:"product-name" bson:"product-name"`
	Group       [][]string         `json:"group" bson:"group"`
}
