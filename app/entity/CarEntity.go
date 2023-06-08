package entity

import "github.com/jinzhu/gorm"

// CarEntity structure for our blog
type CarEntity struct {
	gorm.Model
	Variable1 string `gorm:"column:Variable1" json:"Variable1"`
	Variable2 string `gorm:"column:variable2" json:"Variable2"`
}

func (CarEntity) TableName() string {
	return "Car"
}
