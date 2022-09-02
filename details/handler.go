package details

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"net/http"
)

func GetAll(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {

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
func Get(w http.ResponseWriter, _ *http.Request, params httprouter.Params) {
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

func Create(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
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
