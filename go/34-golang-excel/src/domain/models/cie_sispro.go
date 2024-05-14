package models

import (
	"time"

	"gorm.io/gorm"
)

type Cie struct {
	gorm.Model
	CieId           uint      `gorm:"column:id;not null"`
	Code            string    `gorm:"column:code;not null;"`
	Name            string    `gorm:"column:name;not null;"`
	Description     string    `gorm:"column:description;not null;"`
	IsAvailable     bool      `gorm:"column:is_available;not null;"`
	IsStandartGel   bool      `gorm:"column:is_standart_gel;not null;"`
	IsStandardMSPS  bool      `gorm:"column:is_standart_msps;not null;"`
	AppliesToSex    uint      `gorm:"column:applies_to_sex;not null;"`
	MinimunAge      uint      `gorm:"column:minimun_age;not null;"`
	MaximunAge      uint      `gorm:"column:maximun_age;not null;"`
	MortalityGroup  uint      `gorm:"column:mortality_group;not null;"`
	ExtraV          string    `gorm:"column:extra_v;not null;"`
	Chapter         string    `gorm:"column:chapter;not null;"`
	Subgroup        uint      `gorm:"column:sub_group;not null;"`
	Category        uint      `gorm:"column:category;not null;"`
	Sex             string    `gorm:"column:sex;not null;"`
	UpdateDate      time.Time `gorm:"column:update_date;not null;"`
	IsPublicPrivate *bool     `gorm:"column:is_public_private;"`
}
