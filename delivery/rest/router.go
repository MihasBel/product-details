package rest

import (
	"github.com/gofiber/swagger"
)

func (r *REST) setURLs() {
	r.app.Get("/swagger/*", swagger.HandlerDefault)
	r.app.Get("/", r.index)

	api := r.app.Group("/api")

	v1 := api.Group("/v1", r.isAuthorizedByApikey)
	v1.Get("/details/all", r.getAll)
	v1.Get("/details/one/:id", r.getByID)
	v1.Post("/details/create", r.create)
}
