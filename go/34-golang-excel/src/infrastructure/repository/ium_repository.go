package repository

import (
	"log"

	"github.com/jdsierrab3991/scripts/34-golang-excel/domain/models"
)

func (repo *Repository) SaveIum(data *models.Ium) error {
	resultData, err := repo.existsSisProByCode(&models.Ium{}, data.Code)
	if err != nil {
		return err
	}

	result, isOk := resultData.(*models.Ium)
	if !isOk {
		log.Fatal("NO MAPPER TO IpsNoReps")
	}
	if result.Code == data.Code {
		return nil
	}
	return repo.db.Save(&data).Error
}
