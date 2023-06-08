package repository

import (
	"project-B/app/entity"
	"project-B/app/repo/db"
)

func GetAllModelEntity() ([]entity.ModelEntity, error) {

	var posts []entity.ModelEntity

	dbConnection := db.GetDB()

	dbConnection.Find(&posts)

	return posts, nil

}

func CreateModelEntity(post *entity.ModelEntity) (*entity.ModelEntity, error) {

	dbConnection := db.GetDB()
	dbConnection.Create(&post)
	return post, nil

}
