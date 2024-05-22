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

type Ffm struct {
	gorm.Model
	FfmId                 uint      `gorm:"column:id;not null"`
	Code                  string    `gorm:"column:code;not null;"`
	Name                  string    `gorm:"column:name;not null;"`
	Description           string    `gorm:"column:description;not null;"`
	IsAvailable           bool      `gorm:"column:is_available;not null;"`
	IsStandartGel         bool      `gorm:"column:is_standart_gel;not null;"`
	IsStandardMSPS        bool      `gorm:"column:is_standart_msps;not null;"`
	Level2Group           bool      `gorm:"column:level_2_group"`
	Level2GroupDefinition string    `gorm:"column:level_2_group_definition"`
	Level3Group           string    `gorm:"column:level_3_group"`
	Level3GroupDefinition string    `gorm:"column:level_12_group_definition"`
	ExtraVI               string    `gorm:"column:extra_vi"`
	ExtraVII              string    `gorm:"column:extra_vii"`
	UpdateDate            time.Time `gorm:"column:update_date;not null;"`
}

type GroupService struct {
	gorm.Model
	GroupServiceId uint      `gorm:"column:id;not null"`
	Code           string    `gorm:"column:code;not null;"`
	Name           string    `gorm:"column:name;not null;"`
	IsAvailable    bool      `gorm:"column:is_available;not null;"`
	IsStandartGel  bool      `gorm:"column:is_standart_gel;not null;"`
	IsStandardMSPS bool      `gorm:"column:is_standart_msps;not null;"`
	UpdateDate     time.Time `gorm:"column:update_date;not null;"`
}