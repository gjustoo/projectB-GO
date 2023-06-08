package entity

import "github.com/jinzhu/gorm"

// TrimEntity structure for our blog
type TrimEntity struct {
	gorm.Model
	Variable1 string `gorm:"column:Variable1" json:"Variable1"`
	Variable2 string `gorm:"column:variable2" json:"Variable2"`
}

func (TrimEntity) TableName() string {
	return "Trim"
}
