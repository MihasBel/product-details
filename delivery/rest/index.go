package rest

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func (r *REST) index(c *fiber.Ctx) error {
	return c.Redirect("/swagger/index.html", http.StatusSeeOther)
}
