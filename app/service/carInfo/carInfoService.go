package carInfoService

import (
	"log"
	"project-B/app/model/carInfo"
	carInfoRepo "project-B/app/repo/carInfo"
)

func GetAllCarInfoModel() []carInfo.CarInfoModel {

	posts, err := carInfoRepo.GetAllCarInfoModel()
	if err != nil {
		log.Println("Error retrieving all carAds ", err)
		return nil
	}
	return posts
}

func CreateCarInfoModel(post *carInfo.CarInfoModel) (*carInfo.CarInfoModel, error) {

	post, err := carInfoRepo.CreateCarInfoModel(post)
	if err != nil {
		log.Println("Error creating carAds ", err)
		return nil, err
	}
	return post, nil
}
