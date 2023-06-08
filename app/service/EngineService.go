package service

import (
	"log"
	"project-B/app/entity"
	carInfoRepo "project-B/app/repo"
)

func GetAllEngineEntity() []entity.EngineEntity {

	posts, err := carInfoRepo.GetAllEngineEntity()
	if err != nil {
		log.Println("Error retrieving all carAds ", err)
		return nil
	}
	return posts
}

func CreateEngineEntity(post *entity.EngineEntity) (*entity.EngineEntity, error) {

	post, err := carInfoRepo.CreateEngineEntity(post)
	if err != nil {
		log.Println("Error creating carAds ", err)
		return nil, err
	}
	return post, nil
}
