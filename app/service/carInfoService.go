package service

import (
	"log"
	"project-B/app/entity"
	carInfoRepo "project-B/app/repo"
)

func GetAllCarInfoEntity() []entity.CarInfoEntity {

	posts, err := carInfoRepo.GetAllCarInfoEntity()
	if err != nil {
		log.Println("Error retrieving all carAds ", err)
		return nil
	}
	return posts
}

func CreateCarInfoEntity(post *entity.CarInfoEntity) (*entity.CarInfoEntity, error) {

	post, err := carInfoRepo.CreateCarInfoEntity(post)
	if err != nil {
		log.Println("Error creating carAds ", err)
		return nil, err
	}
	return post, nil
}
