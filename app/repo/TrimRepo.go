package repository

import (
	"project-B/app/entity"
	"project-B/app/repo/db"
)

func GetAllTrimEntity() ([]entity.TrimEntity, error) {

	var posts []entity.TrimEntity

	dbConnection := db.GetDB()

	dbConnection.Find(&posts)

	return posts, nil

}

func CreateTrimEntity(post *entity.TrimEntity) (*entity.TrimEntity, error) {

	dbConnection := db.GetDB()
	dbConnection.Create(&post)
	return post, nil

}
