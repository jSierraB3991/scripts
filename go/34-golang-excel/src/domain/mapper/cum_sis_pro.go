package mapper

import (
	"log"

	"github.com/jdsierrab3991/scripts/34-golang-excel/domain/libs"
	"github.com/jdsierrab3991/scripts/34-golang-excel/domain/models"
)

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
