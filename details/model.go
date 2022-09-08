package details

import (
	"context"
	"github.com/MihasBel/product-details/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Id     primitive.ObjectID `json:"id" bson:"_id"`
// TODO implement map[string]map[string]string insted of [][]string
type Details struct {
	Id          primitive.ObjectID `json:"id" bson:"_id"`
	ProductName string             `json:"product-name" bson:"product-name"`
	Group       [][]string         `json:"group" bson:"group"`
}
type DetailsWithoutId struct {
	ProductName string     `json:"product-name" bson:"product-name"`
	Group       [][]string `json:"group" bson:"group"`
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
func GetById(id primitive.ObjectID) (Details, error) {
	d := Details{}
	res := config.DetailsCollection.FindOne(context.Background(), bson.D{{"_id", id}})
	err := res.Decode(&d)
	if err != nil {
		return d, err
	}
	return d, nil
}
func InsertOne(d Details) (Details, error) {
	d.Id = primitive.NewObjectID()
	_, err := config.DetailsCollection.InsertOne(context.Background(), d)
	if err != nil {
		return Details{}, err
	}
	return d, nil
}
