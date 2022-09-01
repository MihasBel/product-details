package models

import (
	"context"
	"details/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Id     primitive.ObjectID `json:"id" bson:"_id"`

type Details struct {
	Id   primitive.ObjectID `json:"id" bson:"_id"`
	Name string             `json:"name" bson:"name"`
}

func AllDetails() ([]Details, error) {
	var details []Details
	cur, err := config.DetailsCollection.Find(context.Background(), bson.D{})
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.Background())
	for cur.Next(context.Background()) {
		d := Details{}
		err = cur.Decode(&d)
		if err != nil {
			return nil, err
		}
		details = append(details, d)
	}
	if err = cur.Err(); err != nil {
		return nil, err
	}
	return details, nil
}
