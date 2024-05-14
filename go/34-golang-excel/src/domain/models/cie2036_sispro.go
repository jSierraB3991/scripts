package models

import (
	"time"

	"gorm.io/gorm"
)

type Cie2036 struct {
	gorm.Model
	Cie2036Id       uint      `gorm:"column:id;not null"`
	Code            string    `gorm:"column:code;not null;"`
	Name            string    `gorm:"column:name;not null;"`
	Description     string    `gorm:"column:description;not null;"`
	IsAvailable     bool      `gorm:"column:is_available;not null;"`
	IsStandartGel   bool      `gorm:"column:is_standart_gel;not null;"`
	IsStandardMSPS  bool      `gorm:"column:is_standart_msps;not null;"`
	UpdateDate      time.Time `gorm:"column:update_date;not null;"`
	IsPublicPrivate *bool     `gorm:"column:is_public_private;"`
}
