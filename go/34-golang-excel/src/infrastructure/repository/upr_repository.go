package repository

import (
	"log"

	"github.com/jdsierrab3991/scripts/34-golang-excel/domain/models"
)

func (repo *Repository) SaveUPR(data *models.UPR) error {
	resultData, err := repo.existsSisProByCode(&models.UPR{}, data.Code)
	if err != nil {
		return err
	}

	result, isOk := resultData.(*models.UPR)
	if !isOk {
		log.Fatal("NO MAPPER TO UPR")
	}
	if result.Code == data.Code {
		return nil
	}
	return repo.db.Save(&data).Error
}
