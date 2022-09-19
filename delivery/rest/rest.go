package rest

import (
	"github.com/MihasBel/product-details/internal/app"
	"github.com/MihasBel/product-details/internal/rep"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

type REST struct {
	app *fiber.App
	cfg app.Configuration
	d   rep.Detailer
}

func New(config app.Configuration, d rep.Detailer) *REST {
	app := fiber.New()
	rest := REST{
		app: app,
		cfg: config,
		d:   d,
	}
	rest.setURLs()

	return &rest
}
func (r *REST) Start() {
	if err := r.app.Listen(":8080"); err != nil {
		log.Debug().Err(err).Msg("server status")
	}
}
