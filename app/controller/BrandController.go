package controller

import (
	"encoding/json"
	"net/http"
	"project-B/app/entity"
	BrandEntity "project-B/app/entity"
	"project-B/app/service"
)

func GetAllBrandEntity(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var brands []BrandEntity.BrandEntity

	brands = service.GetAllBrandEntity()

	json.NewEncoder(w).Encode(brands)

}

func CreateBrandEntity(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var brand entity.BrandEntity
	_ = json.NewDecoder(r.Body).Decode(&brand)

	service.CreateBrandEntity(&brand)

	json.NewEncoder(w).Encode(brand)

}
