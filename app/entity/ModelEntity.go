package entity

import (
	"time"

	"github.com/jinzhu/gorm"
)

// ModelEntity structure for our blog
type ModelEntity struct {
	gorm.Model
	Name string    `gorm:"column:name" json:"name"`
	Year time.Time `gorm:"column:year" json:"year"`
}

func (ModelEntity) TableName() string {
	return "Model"
}
