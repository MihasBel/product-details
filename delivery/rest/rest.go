package rest

import (
	"context"
	"github.com/MihasBel/product-details/internal/app"
	"github.com/MihasBel/product-details/internal/rep"
	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"net/http"
	"time"
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
func (r *REST) Start(ctx context.Context) error {
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
func (r *REST) Stop(ctx context.Context) error {
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
