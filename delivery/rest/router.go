package rest

import (
	"github.com/gofiber/swagger"
)

func (r *REST) setURLs() {
	r.app.Get("/swagger/*", swagger.HandlerDefault)
	r.app.Get("/", r.index)

	api := r.app.Group("/api")

	v1 := api.Group("/v1", r.isAuth) //TODO rename
	details := v1.Group("/details")
	details.Get("/all", r.getAll)
	details.Get("/one/:id", r.getByID)
	details.Post("/create", r.create) //TODO rename urls
}
