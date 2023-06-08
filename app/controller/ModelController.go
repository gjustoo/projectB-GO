package controller

import (
	"encoding/json"
	"net/http"
	"project-B/app/entity"
	"project-B/app/service"
)

func GetAllModelEntity(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var models []entity.ModelEntity

	models = service.GetAllModelEntity()

	json.NewEncoder(w).Encode(models)

}

func CreateModelEntity(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var model entity.ModelEntity
	_ = json.NewDecoder(r.Body).Decode(&model)

	service.CreateModelEntity(&model)

	json.NewEncoder(w).Encode(model)

}
