package mongoDb

import (
	"context"
	"github.com/MihasBel/product-details/internal/app"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

var DetailsCollection *mongo.Collection
var DB *mongo.Database

func InitDatabase() {

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(app.Config.ConnectionString))
	if err != nil {
		log.Fatal().Err(err).Msg("cannot connect to mongo client")
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot connect to mongo database")
	}
	DB = client.Database(app.Config.Database)

}
