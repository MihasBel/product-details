package apikey

import (
	"github.com/MihasBel/product-details/internal/app"
	"github.com/gofiber/fiber/v2"
)

// IsAuthorizedByApikey checks the API key in the Authorization header
func IsAuthorizedByApikey(c *fiber.Ctx) error {
	auth := c.GetReqHeaders()["Authorization"]
	if auth != app.Config.APIKey {
		return fiber.ErrUnauthorized
	}
	return c.Next()
}
