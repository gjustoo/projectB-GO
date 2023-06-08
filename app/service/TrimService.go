package service

import (
	"log"
	"project-B/app/entity"
	carInfoRepo "project-B/app/repo"
)

func GetAllTrimEntity() []entity.TrimEntity {

	posts, err := carInfoRepo.GetAllTrimEntity()
	if err != nil {
		log.Println("Error retrieving all carAds ", err)
		return nil
	}
	return posts
}

func CreateTrimEntity(post *entity.TrimEntity) (*entity.TrimEntity, error) {

	post, err := carInfoRepo.CreateTrimEntity(post)
	if err != nil {
		log.Println("Error creating carAds ", err)
		return nil, err
	}
	return post, nil
}
