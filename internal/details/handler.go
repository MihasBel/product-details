package details

import (
	"encoding/json"
	"fmt"
	"github.com/MihasBel/product-details/pkg/apikey"
	"github.com/julienschmidt/httprouter"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

// GetAll godoc
// @Summary Retrieves all product details
// @Produce json
// @Success 200 {array} Details
// @Router /details/all [get]
// @Security ApiKeyAuth
func GetAll(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if !apikey.IsAuthorizedByApikey(w, r) {
		http.Error(w, http.StatusText(401), http.StatusUnauthorized)
		return
	}
	ds, err := AllDetails()
	if err != nil {
		log.Error().Err(err).Msg("error while get all details from db")
	}
	dsj, err := json.Marshal(ds)
	if err != nil {
		log.Error().Err(err).Msg("error while marshal all details")
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
	if !apikey.IsAuthorizedByApikey(w, r) {
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
		log.Error().Err(err).Msg("error while reading details id")
	}
	d, err := GetById(id)
	if err != nil {
		log.Error().Err(err).Msg("error while getting one details by id")
	}
	dj, err := json.Marshal(d)
	if err != nil {
		log.Error().Err(err).Msg("error while marshal one details")
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(dj))
}

// Create godoc
// @Summary Creates a new product-details from the received json document
// @Accept json
// @Param request body _withoutId true "product-details schema"
// @Success 200 {object} Details
// @Router /details/create [post]
// @Security ApiKeyAuth
func Create(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if !apikey.IsAuthorizedByApikey(w, r) {
		http.Error(w, http.StatusText(401), http.StatusUnauthorized)
		return
	}
	d := Details{}
	err := json.NewDecoder(r.Body).Decode(&d)
	if err != nil {
		log.Error().Err(err).Msg("error while decode details")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	d, err = InsertOne(d)
	if err != nil {
		log.Error().Err(err).Msg("error while insert one details to db")
	}
	dj, err := json.Marshal(d)
	if err != nil {
		log.Error().Err(err).Msg("error while marshal inserted one details")
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, string(dj))
}
