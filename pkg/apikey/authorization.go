package apikey

import (
	"github.com/MihasBel/product-details/internal/app"
	"github.com/gofiber/fiber/v2"
)

func IsAuthorizedByApikey(c *fiber.Ctx) error {
	auth := c.GetReqHeaders()["Authorization"]
	if auth != app.Config.ApiKey {
		return fiber.ErrUnauthorized
	}
	return c.Next()
}
