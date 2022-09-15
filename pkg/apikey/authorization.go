package apikey

import (
	"github.com/MihasBel/product-details/internal/app"
	"net/http"
)

func IsAuthorizedByApikey(_ http.ResponseWriter, r *http.Request) bool {
	auth := r.Header.Get("Authorization")
	if auth != app.Config.ApiKey {
		return false
	}
	return true
}
