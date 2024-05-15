package models

import (
	"gorm.io/gorm"
)

type CupsRips struct {
	gorm.Model
	CupsRipsId     uint    `gorm:"column:id;not null"`
	Code           string  `gorm:"column:code;not null;"`
	Name           string  `gorm:"column:name;not null;"`
	Description    string  `gorm:"column:description;not null;"`
	IsAvailable    bool    `gorm:"column:is_available;not null;"`
	IsStandartGel  bool    `gorm:"column:is_standart_gel;not null;"`
	IsStandardMSPS bool    `gorm:"column:is_standart_msps;not null;"`
	CupCode        string  `gorm:"column:cup_code;not null;"`
	MaximunNumber  uint    `gorm:"column:maximun_number;not null;"`
	MinimunNumber  uint    `gorm:"column:minimun_number;not null;"`
	Qx             string  `gorm:"column:qx;not null;"`
	DxRequired     string  `gorm:"column:dx_required;not null;"`
	Sex            *string `gorm:"column:sex;"`
	Ambit          *string `gorm:"column:ambit;"`
	Stay           *string `gorm:"column:stay;"`
	Coverage       *string `gorm:"column:coverage;"`
	Duplicate      *string `gorm:"column:duplicate;"`
}
