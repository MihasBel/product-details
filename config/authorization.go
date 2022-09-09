package config

import (
	"net/http"
)

func IsAuthorizedByApikey(_ http.ResponseWriter, r *http.Request) bool {
	auth := r.Header.Get("Authorization")
	if auth != Config.ApiKey {
		return false
	}
	return true
}
