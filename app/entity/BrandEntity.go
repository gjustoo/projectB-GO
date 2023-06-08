package entity

import "github.com/jinzhu/gorm"

// BrandEntity structure for our blog
type BrandEntity struct {
	gorm.Model
	Name string `gorm:"column:name" json:"name"`
	// Country string `gorm:"column:variable2" json:"Variable2"`
}

func (BrandEntity) TableName() string {
	return "Brand"
}
