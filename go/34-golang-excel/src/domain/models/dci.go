package models

import (
	"time"

	"gorm.io/gorm"
)

type Dci struct {
	gorm.Model
	DciId          uint      `gorm:"column:id;not null"`
	Code           string    `gorm:"column:code;not null;"`
	Name           string    `gorm:"column:name;not null;"`
	Extra          *string   `gorm:"column:extra"`
	ExtraVI        *uint     `gorm:"column:extra_vi"`
	IsAvailable    bool      `gorm:"column:is_available;not null;"`
	IsStandartGel  bool      `gorm:"column:is_standart_gel;not null;"`
	IsStandardMSPS bool      `gorm:"column:is_standart_msps;not null;"`
	UpdateDate     time.Time `gorm:"column:update_date;not null;"`
}
