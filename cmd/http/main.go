package main

import (
	"flag"
	"github.com/MihasBel/product-details/internal/app"
	"github.com/jinzhu/configor"

	"net/http"
	"os"

	"github.com/MihasBel/product-details/internal/details"
	"github.com/MihasBel/product-details/pkg/mongoDb"

	_ "github.com/MihasBel/product-details/api/docs"
	"github.com/julienschmidt/httprouter"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	httpSwagger "github.com/swaggo/http-swagger"
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
	mongoDb.InitDatabase()
	details.InitCollection()
}

// @title Details API
// @version 1.0
// @description Swagger API service to store and modify the product details description of any goods

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func main() {
	defer logf.Close()
	router := httprouter.New()
	router.GET("/docs/:any", swaggerHandler)
	router.GET("/", index)

	router.GET("/details/all", details.GetAll)
	router.GET("/details/one/:id", details.Get)
	router.POST("/details/create", details.Create)
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Debug().Err(err).Msg("server status")
	}
}

func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	http.Redirect(w, r, "/docs/index.html", http.StatusSeeOther)
}
func swaggerHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	httpSwagger.WrapHandler(w, r)
}
func logFile() *os.File {
	f, err := os.OpenFile("details.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal().Err(err).Msg("error while opening log file")
	}
	return f
}
