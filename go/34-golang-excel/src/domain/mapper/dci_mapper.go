package mapper

import (
	"log"

	"github.com/jdsierrab3991/scripts/34-golang-excel/domain/libs"
	"github.com/jdsierrab3991/scripts/34-golang-excel/domain/models"
)

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
