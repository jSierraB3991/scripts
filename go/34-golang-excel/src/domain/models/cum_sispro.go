package models

import (
	"time"

	"gorm.io/gorm"
)

type CumSispro struct {
	gorm.Model
	CumSisProId           uint      `gorm:"column:id;not null"`
	Code                  string    `gorm:"column:code;not null;"`
	Name                  string    `gorm:"column:name;not null;"`
	Description           string    `gorm:"column:description;not null;"`
	Available             bool      `gorm:"column:available;not null;"`
	IsStandartGel         bool      `gorm:"column:is_standart_gel;not null;"`
	IsStandardMSPS        bool      `gorm:"column:is_standart_msps;not null;"`
	IndicatorSampleMedic  bool      `gorm:"column:indicator_sample_medic;not null;"`
	AtcCode               string    `gorm:"column:atc_code;not null;"`
	ATC                   string    `gorm:"column:atc;not null;"`
	HealthRegister        string    `gorm:"column:health_register;not null;"`
	ActivePrinciple       string    `gorm:"column:active_principle;not null;"`
	AmountActivePrinciple string    `gorm:"column:amount_active_principle;not null;"`
	ViaAdministratio      string    `gorm:"column:via_administration;not null;"`
	AmountPresentation    uint      `gorm:"column:amount_presentation;not null;"`
	UpdateDate            time.Time `gorm:"column:update_date;not null;"`
}
