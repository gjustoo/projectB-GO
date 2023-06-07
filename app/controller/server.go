package controller

import (
	"fmt"
	"log"
	"net/http"
	carInfoController "project-B/app/controller/carInfo"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

var router *mux.Router

func initCarInfoModelHandlers() {

	router.HandleFunc("/api/carInfo/", carInfoController.GetAllCarInfoModel).Methods("GET")
	router.HandleFunc("/api/carInfo/", carInfoController.CreateCarInfoModel).Methods("POST")
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

	initCarInfoModelHandlers()
	fmt.Printf("router initialized and listening on 3200\n")
	log.Fatal(http.ListenAndServe(":3200", router))
}
