package details

import (
	"context"

	"github.com/MihasBel/product-details/model"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// All fake for test get all documents from DB
func (m MongoDetail) All(ctx context.Context) ([]model.Detail, error) {
	cur, err := m.collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := cur.Close(ctx); err != nil {
			log.Error().Err(err).Msg("error while connecting to details collection")
		}
	}()
	details := make([]model.Detail, 0, cur.RemainingBatchLength())

	if err = cur.All(ctx, &details); err != nil {
		return nil, err
	}
	return details, nil
}

// ByID get one detail by id from DB
func (m MongoDetail) ByID(ctx context.Context, id string) (model.Detail, error) {
	d := model.Detail{}
	idObj, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return d, err
	}
	res := m.collection.FindOne(ctx, bson.D{{Key: "_id", Value: idObj}})
	if err := res.Decode(&d); err != nil {
		return d, err
	}
	return d, nil
}

// InsertOne insert one new detail to DB. Should generate new object id
func (m MongoDetail) InsertOne(ctx context.Context, d model.Detail) (model.Detail, error) {
	dDB := details{
		ID:          primitive.NewObjectID(),
		ProductName: d.ProductName,
		Group:       d.Group,
	}
	_, err := m.collection.InsertOne(ctx, dDB)
	if err != nil {
		return model.Detail{}, err
	}
	d.ID = dDB.ID.Hex()
	return d, nil
}
