package entity

import (
	"time"

	"github.com/jinzhu/gorm"
)

// EngineEntity structure for our blog
type EngineEntity struct {
	gorm.Model
	Code string    `gorm:"column:code" json:"code"`
	Year time.Time `gorm:"column:year" json:"year"`
}

func (EngineEntity) TableName() string {
	return "Engine"
}
