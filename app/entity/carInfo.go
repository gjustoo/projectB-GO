package entity

import "github.com/jinzhu/gorm"

// CarInfoEntity structure for our blog
type CarInfoEntity struct {
	gorm.Model
	Brand     string `gorm:"column:brand" json:"Brand"`
	CarEntity string `gorm:"column:entity" json:"Entity"`
}

func (CarInfoEntity) TableName() string {
	return "car_info"
}
