package carInfoController

import (
	"encoding/json"
	"net/http"
	"project-B/app/model/carInfo"
	carInfoService "project-B/app/service/carInfo"
)

func GetAllCarInfoModel(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var carInfo []carInfo.CarInfoModel

	carInfo = carInfoService.GetAllCarInfoModel()

	json.NewEncoder(w).Encode(carInfo)

}

func CreateCarInfoModel(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var carInfo carInfo.CarInfoModel

	_ = json.NewDecoder(r.Body).Decode(&carInfo)

	carInfoService.CreateCarInfoModel(&carInfo)

	json.NewEncoder(w).Encode(carInfo)

}
