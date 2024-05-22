package repository

import (
	"log"

	"github.com/jdsierrab3991/scripts/34-golang-excel/domain/models"
)

func (repo *Repository) SaveIPSCodHabilitacion(data *models.IpsCpdeHabilitation) error {
	resultData, err := repo.existsSisProByCode(&models.IpsCpdeHabilitation{}, data.Code)
	if err != nil {
		return err
	}

	result, isOk := resultData.(*models.IpsCpdeHabilitation)
	if !isOk {
		log.Fatal("NO MAPPER TO IpsCpdeHabilitation")
	}
	if result.Code == data.Code {
		return nil
	}
	return repo.db.Save(&data).Error
}
