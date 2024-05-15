package models

import "gorm.io/gorm"

type Scrapp struct {
	gorm.Model
	ScrappId uint   `gorm:"column:id;not null"`
	Type     string `gorm:"column:type;not null;"`
	Code     string `gorm:"column:code;not null;"`
}
