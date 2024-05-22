package repository

import (
	"log"

	"github.com/jdsierrab3991/scripts/34-golang-excel/domain/models"
)

func (repo *Repository) SaveMunicipality(data *models.Municipality) error {
	resultData, err := repo.existsSisProByCode(&models.Municipality{}, data.Code)
	if err != nil {
		return err
	}

	result, isOk := resultData.(*models.Municipality)
	if !isOk {
		log.Fatal("NO MAPPER TO Municipality")
	}
	if result.Code == data.Code {
		return nil
	}
	return repo.db.Save(&data).Error
}
