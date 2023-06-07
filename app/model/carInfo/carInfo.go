package carInfo

import "github.com/jinzhu/gorm"

// CarInfoModel structure for our blog
type CarInfoModel struct {
	gorm.Model
	Brand    string `gorm:"column:brand" json:"Brand"`
	CarModel string `gorm:"column:model" json:"Model"`
}

func (CarInfoModel) TableName() string {
	return "car_info"
}
