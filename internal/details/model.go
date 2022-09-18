package details

import (
	"context"

	"github.com/rs/zerolog/log"

	"github.com/MihasBel/product-details/internal/app"
	"github.com/MihasBel/product-details/pkg/mongodb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var collection *mongo.Collection

// InitCollection initializing the details collection
func InitCollection() {
	collection = mongodb.DB.Collection(app.Config.Collection)
}

// Details ID     primitive.ObjectID `json:"id" bson:"_id"`
// TODO implement map[string]map[string]string insted of [][]string
// Details model to represent details in database
type Details struct {
	ID          primitive.ObjectID `json:"id" bson:"_id"`
	ProductName string             `json:"product-name" bson:"product-name"`
	Group       [][]string         `json:"group" bson:"group"`
}

type _withoutID struct {
	ProductName string     `json:"product-name" bson:"product-name"`
	Group       [][]string `json:"group" bson:"group"`
}

func allDetails() ([]Details, error) {
	var details []Details
	cur, err := collection.Find(context.Background(), bson.D{})
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := cur.Close(context.Background()); err != nil {
			log.Error().Err(err).Msg("error while connecting to details collection")
		}
	}()
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
func getByID(id primitive.ObjectID) (Details, error) {
	d := Details{}
	q := bson.D{{Key: "_id", Value: id}}
	res := collection.FindOne(context.Background(), q)
	err := res.Decode(&d)
	if err != nil {
		return d, err
	}
	return d, nil
}
func insertOne(d Details) (Details, error) {
	d.ID = primitive.NewObjectID()
	_, err := collection.InsertOne(context.Background(), d)
	if err != nil {
		return Details{}, err
	}
	return d, nil
}
