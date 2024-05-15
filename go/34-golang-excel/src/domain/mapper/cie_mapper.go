package mapper

import (
	"log"

	"github.com/jdsierrab3991/scripts/34-golang-excel/domain/libs"
	"github.com/jdsierrab3991/scripts/34-golang-excel/domain/models"
)

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
