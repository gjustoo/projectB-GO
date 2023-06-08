package service

import (
	"log"
	"project-B/app/entity"
	carInfoRepo "project-B/app/repo"
)

func GetAllBrandEntity() []entity.BrandEntity {

	posts, err := carInfoRepo.GetAllBrandEntity()
	if err != nil {
		log.Println("Error retrieving all carAds ", err)
		return nil
	}
	return posts
}

func CreateBrandEntity(post *entity.BrandEntity) (*entity.BrandEntity, error) {

	post, err := carInfoRepo.CreateBrandEntity(post)
	if err != nil {
		log.Println("Error creating carAds ", err)
		return nil, err
	}
	return post, nil
}
