package entity

import "github.com/jinzhu/gorm"

// ModelEntity structure for our blog
type ModelEntity struct {
	gorm.Model
	Variable1 string `gorm:"column:Variable1" json:"Variable1"`
	Variable2 string `gorm:"column:variable2" json:"Variable2"`
}

func (ModelEntity) TableName() string {
	return "Model"
}
