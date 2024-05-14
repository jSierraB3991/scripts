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

	}
	return nil
}

func getDatalogoCums(rows []string) *models.CumSispro {
	for _, row := range rows {
		log.Println(row)
	}
	return &models.CumSispro{}
}

func updateSisproFormat(date string) time.Time {
	updateFormat, err := time.Parse(libs.SISPRO_DATE_FORMAT, date)
	if err != nil {
		log.Println("ERROR FORMATING DATE")
		log.Println(err)
		log.Println(date)
		return time.Now()
	}
	return updateFormat
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
		if len(rows) > i && rows[i] != "" {
			return errorsrips.ValidataInMapperSisPro{Data: rows[i]}
		}
	}
	return nil
}
