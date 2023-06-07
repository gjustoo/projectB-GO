package carInfoRepo

import (
	"project-B/app/model/carInfo"
	"project-B/app/repo/db"
)

func GetAllCarInfoModel() ([]carInfo.CarInfoModel, error) {

	var posts []carInfo.CarInfoModel

	dbConnection := db.GetDB()

	dbConnection.Find(&posts)

	return posts, nil

}

func CreateCarInfoModel(post *carInfo.CarInfoModel) (*carInfo.CarInfoModel, error) {

	dbConnection := db.GetDB()
	dbConnection.Create(&post)
	return post, nil

}
