package mongodb

import (
	"context"
	"time"

	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/MihasBel/product-details/internal/app"
)

// DB provides access to the database
var DB *mongo.Database

// Start initializes the connection to the database using the config
func Start() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
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

// Stop closes the connection to the database
func Stop(ctx context.Context) {
	if err := DB.Client().Disconnect(ctx); err != nil {
		log.Error().Err(err).Msg("cannot close connection")
	}
}
