package details

import (
	"details/config"
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"net/http"
)

// GetAll godoc
// @Summary Retrieves all product details
// @Produce json
// @Success 200 {array} Details
// @Router /details/all [get]
// @Security ApiKeyAuth
func GetAll(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if !config.IsAuthorizedByApikey(w, r) {
		http.Error(w, http.StatusText(401), http.StatusUnauthorized)
		return
	}
	ds, err := AllDetails()
	if err != nil {
		log.Println(err)
	}
	dsj, err := json.Marshal(ds)
	if err != nil {
		log.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(dsj))

}

// Get godoc
// @Summary Retrieves product-details on given ID
// @Produce json
// @Param id path string true "product details ID"
// @Success 200 {object} Details
// @Router /details/one/{id} [get]
// @Security ApiKeyAuth
func Get(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	if !config.IsAuthorizedByApikey(w, r) {
		http.Error(w, http.StatusText(401), http.StatusUnauthorized)
		return
	}
	ids := params.ByName("id")
	if !primitive.IsValidObjectID(ids) {
		http.Error(w, "wrong id format "+ids, http.StatusBadRequest)
		return
	}
	id, err := primitive.ObjectIDFromHex(ids)
	if err != nil {
		log.Println(err)
	}
	d, err := DetailsById(id)
	if err != nil {
		log.Println(err)
	}
	dj, err := json.Marshal(d)
	if err != nil {
		log.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(dj))
}

// Create godoc
// @Summary Creates a new product-details from the received json document
// @Accept json
// @Param request body DetailsWithoutId true "product-details schema"
// @Success 200 {object} Details
// @Router /details/create [post]
// @Security ApiKeyAuth
func Create(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if !config.IsAuthorizedByApikey(w, r) {
		http.Error(w, http.StatusText(401), http.StatusUnauthorized)
		return
	}
	d := Details{}
	err := json.NewDecoder(r.Body).Decode(&d)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	d, err = InsertOne(d)
	if err != nil {
		log.Println(err)
	}
	dj, err := json.Marshal(d)
	if err != nil {
		log.Println(dj)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, string(dj))
}
