package mapper

import (
	"time"

	errorsrips "github.com/jdsierrab3991/scripts/34-golang-excel/domain/errors_rips"
	"github.com/jdsierrab3991/scripts/34-golang-excel/domain/libs"
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
