package main

import (
	"github.com/MihasBel/product-details/config"
	"github.com/MihasBel/product-details/details"
	_ "github.com/MihasBel/product-details/docs"
	"github.com/julienschmidt/httprouter"
	httpSwagger "github.com/swaggo/http-swagger"
	"log"
	"net/http"
	"os"
)

var logf *os.File

func init() {
	logf = logFile()
	log.SetOutput(logf)
	config.Config = config.GetConfig()
	config.InitDetailsCollection()
}

// @title Details API
// @version 1.0
// @description Swagger API service to store and modify the product details description of any goods

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func main() {
	defer logf.Close()
	router := httprouter.New()
	router.GET("/docs/:any", swaggerHandler)
	router.GET("/", index)

	router.GET("/details/all", details.GetAll)
	router.GET("/details/one/:id", details.Get)
	router.POST("/details/create", details.Create)
	http.ListenAndServe(":8080", router)
}

func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	http.Redirect(w, r, "/docs/index.html", http.StatusSeeOther)
}
func swaggerHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	httpSwagger.WrapHandler(w, r)
}
func logFile() *os.File {
	f, err := os.OpenFile("details.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	return f
}
