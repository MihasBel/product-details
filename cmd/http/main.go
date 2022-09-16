package main

import (
	"flag"

	"github.com/MihasBel/product-details/internal/app"
	"github.com/MihasBel/product-details/pkg/apikey"
	fiber "github.com/gofiber/fiber/v2"
	"github.com/jinzhu/configor"
	"github.com/julienschmidt/httprouter"
	httpSwagger "github.com/swaggo/http-swagger"

	"net/http"
	"os"

	"github.com/MihasBel/product-details/internal/details"
	"github.com/MihasBel/product-details/pkg/mongoDb"

	_ "github.com/MihasBel/product-details/api/docs"

	"github.com/gofiber/swagger"
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
	mongoDb.InitDatabase()
	details.InitCollection()
}

// @title Details API
// @version 1.0
// @description Swagger API service to store and modify the product details description of any goods
// @BasePath  /api/v1

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	defer logf.Close()
	app := fiber.New()
	app.Get("/swagger/*", swagger.HandlerDefault)
	app.Get("/", index)

	api := app.Group("/api")

	v1 := api.Group("/v1", apikey.IsAuthorizedByApikey)
	v1.Get("/details/all", details.GetAll)
	v1.Get("/details/one/:id", details.Get)
	v1.Post("/details/create", details.Create)

	if err := app.Listen(":8080"); err != nil {
		log.Debug().Err(err).Msg("server status")
	}
}

func index(c *fiber.Ctx) error {
	c.Redirect("/swagger/index.html", http.StatusSeeOther)
	return nil
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
