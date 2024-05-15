package mapper

import (
	"log"

	"github.com/jdsierrab3991/scripts/34-golang-excel/domain/models"
)

func getDataUserEgrese(rows []string) *models.UserEgrese {
	err := validateVoidData(rows, []int{3, 5, 13, 14, 15, 16, 17, 18, 19, 21})
	if err != nil {
		log.Panic(err)
	}
	return &models.UserEgrese{
		Code:            rows[1],
		Name:            rows[2],
		IsAvailable:     rows[4] == "SI",
		IsStandartGel:   rows[6] == "True",
		IsStandardMSPS:  rows[7] == "True",
		Consulation:     rows[8],
		Procedure:       rows[9],
		Emergency:       rows[10] == "SI",
		Hospitalization: rows[11] == "SI",
		Born:            rows[12] == "SI",
		UpdateDate:      *updateSisproFormat(rows[20]),
	}
}
