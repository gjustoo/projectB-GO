package repository

import (
	"project-B/app/entity"
	"project-B/app/repo/db"
)

func GetAllCarInfoEntity() ([]entity.CarInfoEntity, error) {

	var posts []entity.CarInfoEntity

	dbConnection := db.GetDB()

	dbConnection.Find(&posts)

	return posts, nil

}

func CreateCarInfoEntity(post *entity.CarInfoEntity) (*entity.CarInfoEntity, error) {

	dbConnection := db.GetDB()
	dbConnection.Create(&post)
	return post, nil

}
