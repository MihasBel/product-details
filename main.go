package main

import (
	"details/details"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"os"
)

func main() {
	logf := logFile()
	log.SetOutput(logf)
	defer logf.Close()
	router := httprouter.New()
	router.GET("/", index)
	router.GET("/details/all", details.GetAll)
	router.GET("/details/one/:id", details.Get)
	router.POST("/details/create", details.Create)
	http.ListenAndServe(":8080", router)
}
func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
	}
	fmt.Fprint(w, "Hello")
}
func logFile() *os.File {
	f, err := os.OpenFile("details.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	return f
}
