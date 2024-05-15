package models

import (
	"time"

	"gorm.io/gorm"
)

type UserEgrese struct {
	gorm.Model
	UserEgreseId    uint      `gorm:"column:id;not null"`
	Code            string    `gorm:"column:code;not null;"`
	Name            string    `gorm:"column:name;not null;"`
	IsAvailable     bool      `gorm:"column:is_available;not null;"`
	IsStandartGel   bool      `gorm:"column:is_standart_gel;not null;"`
	IsStandardMSPS  bool      `gorm:"column:is_standart_msps;not null;"`
	Consulation     string    `gorm:"column:consulation;not null;"`
	Procedure       string    `gorm:"column:procedure;not null;"`
	Emergency       bool      `gorm:"column:emergency;not null;"`
	Hospitalization bool      `gorm:"column:hospitalization;not null;"`
	Born            bool      `gorm:"column:born;not null;"`
	UpdateDate      time.Time `gorm:"column:update_date;not null;"`
}
