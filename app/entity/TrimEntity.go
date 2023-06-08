package entity

import "github.com/jinzhu/gorm"

// TrimEntity structure for our blog
type TrimEntity struct {
	gorm.Model
	Name string `gorm:"column:name" json:"name"`
	// model string `gorm:"column:model" json:"model"`
}

func (TrimEntity) TableName() string {
	return "Trim"
}
