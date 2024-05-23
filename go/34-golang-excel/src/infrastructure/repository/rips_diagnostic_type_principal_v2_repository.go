package repository

import (
	"log"

	"github.com/jdsierrab3991/scripts/34-golang-excel/domain/models"
)

func (repo *Repository) SaveRipsDiagnostictypePrincipalv2(data *models.RipsDiagnostictypePrincipalv2) error {
	resultData, err := repo.existsSisProByCode(&models.RipsDiagnostictypePrincipalv2{}, data.Code)
	if err != nil {
		return err
	}

	result, isOk := resultData.(*models.RipsDiagnostictypePrincipalv2)
	if !isOk {
		log.Fatal("NO MAPPER TO RipsDiagnostictypePrincipalv2")
	}
	if result.Code == data.Code {
		return nil
	}
	return repo.db.Save(&data).Error
}
