package repository

import (
	"project-B/app/entity"
	"project-B/app/repo/db"
)

func GetAllBrandEntity() ([]entity.BrandEntity, error) {

	var posts []entity.BrandEntity

	dbConnection := db.GetDB()

	dbConnection.Find(&posts)

	return posts, nil

}

func CreateBrandEntity(post *entity.BrandEntity) (*entity.BrandEntity, error) {

	dbConnection := db.GetDB()
	dbConnection.Create(&post)
	return post, nil

}
