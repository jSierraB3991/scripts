package model

import "gorm.io/gorm"

type Municipality struct {
	gorm.Model
	MunicipalityId uint   `gorm:"column:id;not null"`
	Name           string `gorm:"column:name;not null"`
	Code           string `gorm:"column:code;not null"`
	SubRegion      string `gorm:"column:sub_region_code;not null"`
}
