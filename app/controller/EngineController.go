package controller

import (
	"encoding/json"
	"net/http"
	"project-B/app/entity"
	"project-B/app/service"
)

func GetAllEngineEntity(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var engines []entity.EngineEntity

	engines = service.GetAllEngineEntity()

	json.NewEncoder(w).Encode(engines)

}

func CreateEngineEntity(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var engine entity.EngineEntity
	_ = json.NewDecoder(r.Body).Decode(&engine)

	service.CreateEngineEntity(&engine)

	json.NewEncoder(w).Encode(engine)

}
