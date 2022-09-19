package rest

import (
	"github.com/MihasBel/product-details/internal/app"
	"github.com/gofiber/fiber/v2"
)

func (r *REST) isAuthorizedByApikey(c *fiber.Ctx) error {
	auth := c.GetReqHeaders()["Authorization"]
	if auth != app.Config.APIKey {
		return fiber.ErrUnauthorized
	}
	return c.Next()
}
