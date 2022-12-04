package main

import (
	"context"
	"flag"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/MihasBel/product-details/internal/mgdetailer"

	"github.com/MihasBel/product-details/delivery/rest"

	"github.com/MihasBel/product-details/internal/app"
	"github.com/jinzhu/configor"

	"github.com/MihasBel/product-details/pkg/mongodb"

	_ "github.com/MihasBel/product-details/api/docs"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var logf *os.File

var configPath string

func init() {
	log.Logger = log.Output(zerolog.New(logFile()))
	flag.StringVar(&configPath, "config", "configs/local/env.json", "Config file path")
	flag.Parse()

	if err := configor.New(&configor.Config{ErrorOnUnmatchedKeys: true}).Load(&app.Config, configPath); err != nil {
		log.Fatal().Err(err).Msg("cannot load config")
	}
	mongodb.Start()
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
	cfg := app.Config
	detailer := mgdetailer.New(mongodb.DB.Collection(cfg.Collection))
	application := rest.New(cfg, detailer)

	startCtx, startCancel := context.WithTimeout(context.Background(), time.Duration(cfg.StartTimeout)*time.Second)
	defer startCancel()
	if err := application.Start(startCtx); err != nil {
		log.Fatal().Err(err).Msg("cannot start application") // nolint
	}

	log.Info().Msg("application started")

	quitCh := make(chan os.Signal, 1)
	signal.Notify(quitCh, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-quitCh

	stopCtx, stopCancel := context.WithTimeout(context.Background(), time.Duration(cfg.StartTimeout)*time.Second)
	defer stopCancel()

	if err := application.Stop(stopCtx); err != nil {
		log.Error().Err(err).Msg("cannot stop application")
	}
	mongodb.Stop(stopCtx)
	log.Info().Msg("service is down")
}

func logFile() *os.File {
	f, err := os.OpenFile("details.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal().Err(err).Msg("error while opening log file")
	}
	return f
}
