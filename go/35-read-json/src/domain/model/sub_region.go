package model

import "gorm.io/gorm"

type SubRegion struct {
	gorm.Model
	SubRegionId uint   `gorm:"column:id;not null"`
	Name        string `gorm:"column:name;not null"`
	Region      string `gorm:"column:Region;not null"`
	Code        string `gorm:"column:code;not null"`
}
