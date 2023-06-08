package controller

import (
	"encoding/json"
	"net/http"
	"project-B/app/entity"
	"project-B/app/service"
)

func GetAllTrimEntity(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var trims []entity.TrimEntity

	trims = service.GetAllTrimEntity()

	json.NewEncoder(w).Encode(trims)

}

func CreateTrimEntity(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var trim entity.TrimEntity
	_ = json.NewDecoder(r.Body).Decode(&trim)

	service.CreateTrimEntity(&trim)

	json.NewEncoder(w).Encode(trim)

}
