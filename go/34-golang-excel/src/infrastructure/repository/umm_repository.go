package repository

import (
	"log"

	"github.com/jdsierrab3991/scripts/34-golang-excel/domain/models"
)

func (repo *Repository) SaveUMM(data *models.UMM) error {
	resultData, err := repo.existsSisProByCode(&models.UMM{}, data.Code)
	if err != nil {
		return err
	}

	result, isOk := resultData.(*models.UMM)
	if !isOk {
		log.Fatal("NO MAPPER TO UMM")
	}
	if result.Code == data.Code {
		return nil
	}
	return repo.db.Save(&data).Error
}
