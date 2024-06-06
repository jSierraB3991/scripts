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

type IpsCpdeHabilitation struct {
	gorm.Model
	IpsCpdeHabilitationId uint      `gorm:"column:id;not null"`
	Code                  string    `gorm:"column:code;not null;"`
	Name                  string    `gorm:"column:name;not null;"`
	IsAvailable           bool      `gorm:"column:is_available;not null;"`
	IsStandartGel         bool      `gorm:"column:is_standart_gel;not null;"`
	IsStandardMSPS        bool      `gorm:"column:is_standart_msps;not null;"`
	TypeIdPres            string    `gorm:"column:type_id_pres;not null;"`
	NumIdPres             string    `gorm:"column:num_id_pres;not null;"`
	CodePres              string    `gorm:"column:code_pres;not null;"`
	CodeMpiSede           string    `gorm:"column:code_mpi_sede;not null;"`
	NameMpiSede           string    `gorm:"column:name_mpi_sede;not null;"`
	NameDeptoSede         string    `gorm:"column:name_depto_sede;not null;"`
	ClassPres             uint      `gorm:"column:class_pres;not null;"`
	NameClassPres         string    `gorm:"column:name_class_pres;not null;"`
	UpdateDate            time.Time `gorm:"column:update_date;not null;"`
}

type IpsNoReps struct {
	gorm.Model
	IpsNoRepsId      uint      `gorm:"column:id;not null"`
	Code             string    `gorm:"column:code;not null;"`
	Name             string    `gorm:"column:name;not null;"`
	Description      string    `gorm:"column:description;not null;"`
	IsAvailable      bool      `gorm:"column:is_available;not null;"`
	IsStandartGel    bool      `gorm:"column:is_standart_gel;not null;"`
	IsStandardMSPS   bool      `gorm:"column:is_standart_msps;not null;"`
	Telphone         string    `gorm:"column:telphone;not null;"`
	Manager          string    `gorm:"column:manager;not null;"`
	Regime           string    `gorm:"column:regime;not null;"`
	CodeDepto        string    `gorm:"column:code_depto;not null;"`
	Department       string    `gorm:"column:department;not null;"`
	CodeMunicipality string    `gorm:"column:code_municipality;not null;"`
	Municipality     string    `gorm:"column:municipality;not null;"`
	IpsType          string    `gorm:"column:ips_type;not null;"`
	AtentionLevel    uint      `gorm:"column:atention_level;not null;"`
	Nit              string    `gorm:"column:nit;not null;"`
	UpdateDate       time.Time `gorm:"column:update_date;not null;"`
	IsPublicPrivate  *bool     `gorm:"column:is_public_private;"`
}

type Ium struct {
	gorm.Model
	IumId                        uint      `gorm:"column:id;not null"`
	Code                         string    `gorm:"column:code;not null;"`
	Name                         string    `gorm:"column:name;not null;"`
	IsAvailable                  bool      `gorm:"column:is_available;not null;"`
	IsStandartGel                bool      `gorm:"column:is_standart_gel;not null;"`
	IsStandardMSPS               bool      `gorm:"column:is_standart_msps;not null;"`
	Nivel1                       string    `gorm:"column:level_one;not null;"`
	ActivePrincipal              string    `gorm:"column:active_principal;not null;"`
	CodeActivePrincipal          string    `gorm:"column:code_active_principal;not null;"`
	PharmaceuticalForm           string    `gorm:"column:pharmaceutical_form;not null;"`
	CodePharmaceuticalForm       string    `gorm:"column:code_pharmaceutical_form;not null;"`
	Nivel2                       uint      `gorm:"column:level_two;not null;"`
	CodeComercialitionForm       string    `gorm:"column:code_comercialization_form;not null;"`
	Nivel3                       uint      `gorm:"column:level_three;not null;"`
	ConditionResgiterMedicSample string    `gorm:"column:condition_register_medic_sample;not null;"`
	PackegeUnique                string    `gorm:"column:package_uniq;not null;"`
	UpdateDate                   time.Time `gorm:"column:update_date;not null;"`
}

type AtentionModality struct {
	gorm.Model
	AtentionModalityId uint      `gorm:"column:id;not null"`
	Code               string    `gorm:"column:code;not null;"`
	Name               string    `gorm:"column:name;not null;"`
	IsAvailable        bool      `gorm:"column:is_available;not null;"`
	IsStandartGel      bool      `gorm:"column:is_standart_gel;not null;"`
	IsStandardMSPS     bool      `gorm:"column:is_standart_msps;not null;"`
	UpdateDate         time.Time `gorm:"column:update_date;not null;"`
}

type RipsCausaExternaV2 struct {
	gorm.Model
	RipsCausaExternaV2Id uint      `gorm:"column:id;not null"`
	Code                 string    `gorm:"column:code;not null;"`
	Name                 string    `gorm:"column:name;not null;"`
	IsAvailable          bool      `gorm:"column:is_available;not null;"`
	IsStandartGel        bool      `gorm:"column:is_standart_gel;not null;"`
	IsStandardMSPS       bool      `gorm:"column:is_standart_msps;not null;"`
	Consults             string    `gorm:"column:consults;not null;"`
	Procedure            string    `gorm:"column:procedure;not null;"`
	Urgency              string    `gorm:"column:urgency;not null;"`
	Hospitalization      string    `gorm:"column:hospitalization;not null;"`
	RnBorn               string    `gorm:"column:rn_born;not null;"`
	UpdateDate           time.Time `gorm:"column:update_date;not null;"`
}
type RipsConsultFinalV2 struct {
	gorm.Model
	RipsConsultFinalV2Id uint      `gorm:"column:id;not null"`
	Code                 string    `gorm:"column:code;not null;"`
	Name                 string    `gorm:"column:name;not null;"`
	IsAvailable          bool      `gorm:"column:is_available;not null;"`
	IsStandartGel        bool      `gorm:"column:is_standart_gel;not null;"`
	IsStandardMSPS       bool      `gorm:"column:is_standart_msps;not null;"`
	Consults             string    `gorm:"column:consults;not null;"`
	Procedure            string    `gorm:"column:procedure;not null;"`
	Urgency              string    `gorm:"column:urgency;not null;"`
	Hospitalization      string    `gorm:"column:hospitalization;not null;"`
	RnBorn               string    `gorm:"column:rn_born;not null;"`
	UpdateDate           time.Time `gorm:"column:update_date;not null;"`
}
type RipsDiagnostictypePrincipalv2 struct {
	gorm.Model
	RipsDiagnostictypePrincipalv2Id uint      `gorm:"column:id;not null"`
	Code                            string    `gorm:"column:code;not null;"`
	Name                            string    `gorm:"column:name;not null;"`
	IsAvailable                     bool      `gorm:"column:is_available;not null;"`
	IsStandartGel                   bool      `gorm:"column:is_standart_gel;not null;"`
	IsStandardMSPS                  bool      `gorm:"column:is_standart_msps;not null;"`
	UpdateDate                      time.Time `gorm:"column:update_date;not null;"`
}

type UserType struct {
	gorm.Model
	UserTypeId     uint      `gorm:"column:id;not null"`
	Code           string    `gorm:"column:code;not null;"`
	Name           string    `gorm:"column:name;not null;"`
	IsAvailable    bool      `gorm:"column:is_available;not null;"`
	IsStandartGel  bool      `gorm:"column:is_standart_gel;not null;"`
	IsStandardMSPS bool      `gorm:"column:is_standart_msps;not null;"`
	UpdateDate     time.Time `gorm:"column:update_date;not null;"`
}
type Service struct {
	gorm.Model
	ServiceId      uint      `gorm:"column:id;not null"`
	Code           string    `gorm:"column:code;not null;"`
	Description    string    `gorm:"column:description;not null;"`
	Name           string    `gorm:"column:name;not null;"`
	IsAvailable    bool      `gorm:"column:is_available;not null;"`
	IsStandartGel  bool      `gorm:"column:is_standart_gel;not null;"`
	IsStandardMSPS bool      `gorm:"column:is_standart_msps;not null;"`
	UpdateDate     time.Time `gorm:"column:update_date;not null;"`

	Extra_I    string `gorm:"column:extra_i;not null;"`
	Extra_II   string `gorm:"column:extra_ii;not null;"`
	Extra_III  string `gorm:"column:extra_iii;not null;"`
	Extra_IV   string `gorm:"column:extra_iv;not null;"`
	Extra_V    string `gorm:"column:extra_v;not null;"`
	Extra_VI   string `gorm:"column:extra_vi;not null;"`
	Extra_VII  string `gorm:"column:extra_vii;not null;"`
	Extra_VIII string `gorm:"column:extra_viii;not null;"`
	Extra_IX   string `gorm:"column:extra_ix;not null;"`
	Extra_X    string `gorm:"column:extra_x;not null;"`
}
type OtherService struct {
	gorm.Model
	OtherServiceId uint      `gorm:"column:id;not null"`
	Code           string    `gorm:"column:code;not null;"`
	Name           string    `gorm:"column:name;not null;"`
	IsAvailable    bool      `gorm:"column:is_available;not null;"`
	IsStandartGel  bool      `gorm:"column:is_standart_gel;not null;"`
	IsStandardMSPS bool      `gorm:"column:is_standart_msps;not null;"`
	UpdateDate     time.Time `gorm:"column:update_date;not null;"`
}
type TypeIdPISIS struct {
	gorm.Model
	TypeIdPISISId   uint      `gorm:"column:id;not null"`
	Code            string    `gorm:"column:code;not null;"`
	Name            string    `gorm:"column:name;not null;"`
	IsAvailable     bool      `gorm:"column:is_available;not null;"`
	IsStandartGel   bool      `gorm:"column:is_standart_gel;not null;"`
	ExtraI          string    `gorm:"column:extre_i;not null;"`
	ExtraII         string    `gorm:"column:extre_ii;not null;"`
	ExtraIII        string    `gorm:"column:extre_iii;not null;"`
	ExtraIV         string    `gorm:"column:extra_iv;not null;"`
	ExtraV          string    `gorm:"column:extra_v;not null;"`
	ExtraVI         string    `gorm:"column:extra_vi;not null;"`
	ExtraIX         string    `gorm:"column:extra_ix;not null;"`
	IsStandardMSPS  bool      `gorm:"column:is_standart_msps;not null;"`
	UpdateDate      time.Time `gorm:"column:update_date;not null;"`
	IsPublicPrivate *bool     `gorm:"column:is_public_private;"`
}
type MedicTypePOS struct {
	gorm.Model
	MedicTypePOSId uint      `gorm:"column:id;not null"`
	Code           string    `gorm:"column:code;not null;"`
	Name           string    `gorm:"column:name;not null;"`
	IsAvailable    bool      `gorm:"column:is_available;not null;"`
	IsStandartGel  bool      `gorm:"column:is_standart_gel;not null;"`
	IsStandardMSPS bool      `gorm:"column:is_standart_msps;not null;"`
	UpdateDate     time.Time `gorm:"column:update_date;not null;"`
}
type TypeNote struct {
	gorm.Model
	TypeNoteId     uint      `gorm:"column:id;not null"`
	Code           string    `gorm:"column:code;not null;"`
	Name           string    `gorm:"column:name;not null;"`
	IsAvailable    bool      `gorm:"column:is_available;not null;"`
	IsStandartGel  bool      `gorm:"column:is_standart_gel;not null;"`
	IsStandardMSPS bool      `gorm:"column:is_standart_msps;not null;"`
	UpdateDate     time.Time `gorm:"column:update_date;not null;"`
}
type UMM struct {
	gorm.Model
	UMMId           uint      `gorm:"column:id;not null"`
	Code            string    `gorm:"column:code;not null;"`
	Name            string    `gorm:"column:name;not null;"`
	Description     string    `gorm:"column:description;not null;"`
	IsAvailable     bool      `gorm:"column:is_available;not null;"`
	ExtraI          string    `gorm:"column:extre_i;not null;"`
	ExtraII         string    `gorm:"column:extre_ii;not null;"`
	ExtraIII        string    `gorm:"column:extre_iii;not null;"`
	IsStandartGel   bool      `gorm:"column:is_standart_gel;not null;"`
	IsStandardMSPS  bool      `gorm:"column:is_standart_msps;not null;"`
	UpdateDate      time.Time `gorm:"column:update_date;not null;"`
	IsPublicPrivate *bool     `gorm:"column:is_public_private;"`
}
type UPR struct {
	gorm.Model
	UPRId           uint      `gorm:"column:id;not null"`
	Code            string    `gorm:"column:code;not null;"`
	Name            string    `gorm:"column:name;not null;"`
	IsAvailable     bool      `gorm:"column:is_available;not null;"`
	IsStandartGel   bool      `gorm:"column:is_standart_gel;not null;"`
	IsStandardMSPS  bool      `gorm:"column:is_standart_msps;not null;"`
	UpdateDate      time.Time `gorm:"column:update_date;not null;"`
	IsPublicPrivate *bool     `gorm:"column:is_public_private;"`
}
type IngressUser struct {
	gorm.Model
	IngressUserId   uint      `gorm:"column:id;not null"`
	Code            string    `gorm:"column:code;not null;"`
	Name            string    `gorm:"column:name;not null;"`
	IsAvailable     bool      `gorm:"column:is_available;not null;"`
	IsStandartGel   bool      `gorm:"column:is_standart_gel;not null;"`
	IsStandardMSPS  bool      `gorm:"column:is_standart_msps;not null;"`
	Consult         string    `gorm:"column:consult;not null;"`
	Procedure       string    `gorm:"column:procedure;not null;"`
	Emergency       string    `gorm:"column:emergency;not null;"`
	Hospitalization string    `gorm:"column:hospitalization;not null;"`
	RBorn           string    `gorm:"column:r_born;not null;"`
	UpdateDate      time.Time `gorm:"column:update_date;not null;"`
}
