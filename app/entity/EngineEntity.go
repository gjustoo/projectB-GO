package entity

import "github.com/jinzhu/gorm"

// EngineEntity structure for our blog
type EngineEntity struct {
	gorm.Model
	Variable1 string `gorm:"column:Variable1" json:"Variable1"`
	Variable2 string `gorm:"column:variable2" json:"Variable2"`
}

func (EngineEntity) TableName() string {
	return "Engine"
}
