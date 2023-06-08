package service

import (
	"log"
	"project-B/app/entity"
	carInfoRepo "project-B/app/repo"
)

func GetAllModelEntity() []entity.ModelEntity {

	posts, err := carInfoRepo.GetAllModelEntity()
	if err != nil {
		log.Println("Error retrieving all carAds ", err)
		return nil
	}
	return posts
}

func CreateModelEntity(post *entity.ModelEntity) (*entity.ModelEntity, error) {

	post, err := carInfoRepo.CreateModelEntity(post)
	if err != nil {
		log.Println("Error creating carAds ", err)
		return nil, err
	}
	return post, nil
}
