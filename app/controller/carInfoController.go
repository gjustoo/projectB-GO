package controller

import (
	"encoding/json"
	"net/http"
	carInfoEntity "project-B/app/entity"
	carInfoService "project-B/app/service"
)

func GetAllCarInfoEntity(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var carInfo []carInfoEntity.CarInfoEntity

	carInfo = carInfoService.GetAllCarInfoEntity()

	json.NewEncoder(w).Encode(carInfo)

}

func CreateCarInfoEntity(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var carInfo carInfoEntity.CarInfoEntity
	_ = json.NewDecoder(r.Body).Decode(&carInfo)

	carInfoService.CreateCarInfoEntity(&carInfo)

	json.NewEncoder(w).Encode(carInfo)

}
