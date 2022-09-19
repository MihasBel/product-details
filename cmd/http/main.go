package main

import (
	"flag"
	"github.com/MihasBel/product-details/delivery/rest"
	"os"

	"github.com/MihasBel/product-details/internal/app"
	"github.com/jinzhu/configor"

	"github.com/MihasBel/product-details/internal/details"
	"github.com/MihasBel/product-details/pkg/mongodb"

	_ "github.com/MihasBel/product-details/api/docs"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var logf *os.File
var configPath string

func init() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: logFile()})
	flag.StringVar(&configPath, "config", "env.json", "Config file path")
	flag.Parse()

	if err := configor.New(&configor.Config{ErrorOnUnmatchedKeys: true}).Load(&app.Config, configPath); err != nil {
		log.Fatal().Err(err).Msg("cannot load config")
	}
	mongodb.InitDatabase()
}

// @title Details API
// @version 1.0
// @description Swagger API service to store and modify the product details description of any goods
// @BasePath  /api/v1

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	defer func() {
		if err := logf.Close(); err != nil {
			log.Error().Err(err).Msg("error while closing log file")
		}
	}()
	detailer := details.New(mongodb.DB.Collection(app.Config.Collection))
	app := rest.New(app.Config, detailer)
	app.Start()
}

func logFile() *os.File {
	f, err := os.OpenFile("details.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal().Err(err).Msg("error while opening log file")
	}
	return f
}
