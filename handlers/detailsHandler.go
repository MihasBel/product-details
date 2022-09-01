package handlers

import (
	"details/models"
	"encoding/json"
	"fmt"
	"net/http"
)

func GetAll(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}
	d, err := models.AllDetails()
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotImplemented)
		return
	}
	dj, err := json.Marshal(d)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotImplemented)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(dj))

}
