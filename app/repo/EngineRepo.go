package repository

import (
	"project-B/app/entity"
	"project-B/app/repo/db"
)

func GetAllEngineEntity() ([]entity.EngineEntity, error) {

	var posts []entity.EngineEntity

	dbConnection := db.GetDB()

	dbConnection.Find(&posts)

	return posts, nil

}

func CreateEngineEntity(post *entity.EngineEntity) (*entity.EngineEntity, error) {

	dbConnection := db.GetDB()
	dbConnection.Create(&post)
	return post, nil

}
