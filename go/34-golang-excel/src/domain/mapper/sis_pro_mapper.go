package mapper

import (
	"log"
	"time"

	errorsrips "github.com/jdsierrab3991/scripts/34-golang-excel/domain/errors_rips"
	"github.com/jdsierrab3991/scripts/34-golang-excel/domain/libs"
	"github.com/jdsierrab3991/scripts/34-golang-excel/domain/models"
)

func GetDataSispro(row []string, code string) interface{} {
	switch code {
	case libs.CatalogoCUMs:
		return getDatalogoCums(row)
	case libs.CIE10:
		return getDataCie10(row)
	case libs.CIE10Clasificacion2036:
		return getDataCie102036(row)
	case libs.CUPSRips:
		return getDataCupRips(row)
	case libs.CondicionyDestinoUsuarioEgreso:
		return getDataUserEgrese(row)
	case libs.DCI:
		return getDataDci(row)
	case libs.FFM:
		return getDataFfm(row)
	case libs.GrupoServicios:
		return getGroupService(row)
	}
	return nil
}
func updateSisproFormat(date string) *time.Time {
	updateFormat, err := time.Parse(libs.SISPRO_DATE_FORMAT, date)
	if err != nil {
		return nil
	}
	return &updateFormat
}

func isPublicPrivate(row []string, index int) *bool {
	if len(row) > index {
		result := row[index] == "False"
		return &result
	}
	return nil
}

func validateVoidData(rows []string, indexs []int) error {
	for _, i := range indexs {
		if len(rows) > i && rows[i] != "" && rows[i] != "NULL" {
			return errorsrips.ValidataInMapperSisPro{Data: rows[i]}
		}
	}
	return nil
}

func getPointerString(data string) *string {
	if data == "NULL" {
		return nil
	}
	return &data
}

func getDataCie10(rows []string) *models.Cie {
	err := validateVoidData(rows, []int{5, 14, 18, 19})
	if err != nil {
		log.Panic(err)
	}

	return &models.Cie{
		Code:            rows[1],
		Name:            rows[2],
		Description:     rows[3],
		IsAvailable:     rows[4] == "SI",
		IsStandartGel:   rows[6] == "True",
		IsStandardMSPS:  rows[7] == "True",
		AppliesToSex:    libs.GetUintForString(rows[8]),
		MinimunAge:      libs.GetUintForString(rows[9]),
		MaximunAge:      libs.GetUintForString(rows[10]),
		MortalityGroup:  libs.GetUintForString(rows[11]),
		ExtraV:          rows[12],
		Chapter:         rows[13],
		Subgroup:        libs.GetUintForString(rows[15]),
		Category:        libs.GetUintForString(rows[16]),
		Sex:             rows[17],
		UpdateDate:      *updateSisproFormat(rows[20]),
		IsPublicPrivate: isPublicPrivate(rows, 21),
	}
}

func getDataCie102036(rows []string) *models.Cie2036 {
	err := validateVoidData(rows, []int{5, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19})
	if err != nil {
		log.Panic(err)
	}
	return &models.Cie2036{
		Code:            rows[1],
		Name:            rows[2],
		Description:     rows[3],
		IsAvailable:     rows[4] == "SI",
		IsStandartGel:   rows[6] == "True",
		IsStandardMSPS:  rows[7] == "True",
		UpdateDate:      *updateSisproFormat(rows[20]),
		IsPublicPrivate: isPublicPrivate(rows, 21),
	}
}

func getDatalogoCums(rows []string) *models.CumSispro {
	err := validateVoidData(rows, []int{5, 17, 18, 19, 21})
	if err != nil {
		log.Panic(err)
	}
	return &models.CumSispro{
		Code:                  rows[1],
		Name:                  rows[2],
		Description:           rows[3],
		Available:             rows[4] == "SI",
		IsStandartGel:         rows[6] == "True",
		IsStandardMSPS:        rows[7] == "True",
		IndicatorSampleMedic:  rows[8] == "SI",
		AtcCode:               rows[9],
		ATC:                   rows[10],
		HealthRegister:        rows[11],
		ActivePrinciple:       rows[12],
		AmountActivePrinciple: rows[13],
		ViaAdministratio:      rows[15],
		AmountPresentation:    libs.GetUintForString(rows[16]),
		UpdateDate:            *updateSisproFormat(rows[20]),
	}
}

func getDataCupRips(rows []string) *models.CupsRips {
	err := validateVoidData(rows, []int{5, 18, 19, 21})
	if err != nil {
		log.Panic(err)
	}

	return &models.CupsRips{
		Code:           rows[1],
		Name:           rows[2],
		Description:    rows[3],
		IsAvailable:    rows[4] == "SI",
		IsStandartGel:  rows[6] == "True",
		IsStandardMSPS: rows[7] == "True",
		CupCode:        rows[8],
		Qx:             rows[9],
		MinimunNumber:  libs.GetUintForString(rows[10]),
		MaximunNumber:  libs.GetUintForString(rows[11]),
		DxRequired:     rows[12],
		Sex:            getPointerString(rows[13]),
		Ambit:          getPointerString(rows[14]),
		Stay:           getPointerString(rows[15]),
		Coverage:       getPointerString(rows[16]),
		Duplicate:      getPointerString(rows[17]),
	}
}

func getDataDci(rows []string) *models.Dci {

	err := validateVoidData(rows, []int{3, 5, 9, 10, 11, 12, 14, 15, 16, 17, 18, 19, 21})
	if err != nil {
		log.Panic(err)
	}
	return &models.Dci{
		Code:           rows[1],
		Name:           rows[2],
		IsAvailable:    rows[4] == "SI",
		IsStandartGel:  rows[6] == "True",
		IsStandardMSPS: rows[7] == "True",
		Extra:          libs.GetStringPoint(rows[8]),
		ExtraVI:        libs.GetUintPoint(rows[13]),
		UpdateDate:     *updateSisproFormat(rows[20]),
	}
}

func getDataFfm(rows []string) *models.Ffm {
	err := validateVoidData(rows, []int{5, 8, 16, 17, 18, 19, 21})
	if err != nil {
		log.Panic(err)
	}
	return &models.Ffm{
		Code:                  rows[1],
		Name:                  rows[2],
		Description:           rows[3],
		IsAvailable:           rows[4] == "SI",
		IsStandartGel:         rows[6] == "True",
		IsStandardMSPS:        rows[7] == "True",
		Level2Group:           rows[9] == "VERDADERO",
		Level2GroupDefinition: rows[10],
		Level3Group:           rows[11],
		Level3GroupDefinition: rows[12],
		ExtraVI:               rows[13],
		ExtraVII:              rows[14],
		UpdateDate:            *updateSisproFormat(rows[20]),
	}
}

func getGroupService(rows []string) *models.GroupService {
	err := validateVoidData(rows, []int{3, 5, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 21})
	if err != nil {
		log.Panic(err)
	}
	return &models.GroupService{
		Code:           rows[1],
		Name:           rows[2],
		IsAvailable:    rows[4] == "SI",
		IsStandartGel:  rows[6] == "SI",
		IsStandardMSPS: rows[7] == "SI",
		UpdateDate:     *updateSisproFormat(rows[20]),
	}
}
