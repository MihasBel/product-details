package rest

import (
	"context"
	"net/http"
	"time"

	"github.com/MihasBel/product-details/internal/app"
	"github.com/MihasBel/product-details/internal/rep"
	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

// REST represent REST-full application
type REST struct {
	app *fiber.App
	cfg app.Configuration
	d   rep.Detailer
}

// New Create new instance of REST. Should use only in main.
func New(config app.Configuration, d rep.Detailer) *REST {
	a := fiber.New() //TODO add timeout
	rest := REST{
		app: a,
		cfg: config,
		d:   d,
	}
	rest.setURLs()

	return &rest
}

// Start an application
func (r *REST) Start(_ context.Context) error {
	errCh := make(chan error)
	log.Debug().Msgf("start listening %q", r.cfg.Address)
	go func() {
		if err := r.app.Listen(r.cfg.Address); err != nil && err != http.ErrServerClosed {
			errCh <- errors.Wrap(err, "cannot listen and serve")
		}
	}()

	select {
	case err := <-errCh:
		return err
	case <-time.After(time.Duration(r.cfg.StartTimeout) * time.Second):
		return nil
	}
}

// Stop an application
func (r *REST) Stop(_ context.Context) error {
	errCh := make(chan error)
	log.Debug().Msgf("stopping %q", r.cfg.Address)
	go func() {
		if err := r.app.Shutdown(); err != nil {
			errCh <- errors.Wrap(err, "cannot shutdown")
		}
	}()

	select {
	case err := <-errCh:
		return err
	case <-time.After(time.Duration(r.cfg.StopTimeout) * time.Second):
		return nil

	}
}
