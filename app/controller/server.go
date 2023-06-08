package controller

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

var router *mux.Router

func initCarInfoEntityHandlers() {

	router.HandleFunc("/api/carInfo/", GetAllCarInfoEntity).Methods("GET")
	router.HandleFunc("/api/carInfo/", CreateCarInfoEntity).Methods("POST")
	// router.HandleFunc("/api/carInfo/search", carInfoController.GetAllcarInfosByTitle).Methods("GET")
	// router.HandleFunc("/api/carInfo/{id}", carInfoController.GetcarInfo).Methods("GET")
	// router.HandleFunc("/api/carInfo/{id}", carInfoController.DeletecarInfo).Methods("DELETE")
}

func Start() {
	router = mux.NewRouter()
	router.Use(handlers.CORS(handlers.AllowedOrigins([]string{"*"})))
	router.Use(handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})))
	router.Use(handlers.CORS(handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "DELETE", "OPTIONS"})))
	// Enable CORS

	initCarInfoEntityHandlers()
	fmt.Printf("router initialized and listening on 3200\n")
	log.Fatal(http.ListenAndServe(":3200", router))
}
