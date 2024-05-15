package mapper

import (
	"log"

	"github.com/jdsierrab3991/scripts/34-golang-excel/domain/libs"
	"github.com/jdsierrab3991/scripts/34-golang-excel/domain/models"
)

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
