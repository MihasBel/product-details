package config

import (
	"log"
	"net/http"
)

func IsAuthorizedByApikey(_ http.ResponseWriter, r *http.Request) bool {
	auth := r.Header.Get("Authorization")
	log.Println("apikey - ", auth)
	log.Println("config key", Config.ApiKey)
	if auth != Config.ApiKey {
		return false
	}
	return true
}
