package rest

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (r *REST) index(c *fiber.Ctx) error {
	return c.Redirect("/swagger/index.html", http.StatusSeeOther)
}
