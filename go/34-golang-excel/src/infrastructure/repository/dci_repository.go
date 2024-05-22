package repository

import (
	"log"

	"github.com/jdsierrab3991/scripts/34-golang-excel/domain/models"
)

func (repo *Repository) SaveDci(data *models.Dci) error {
	resultData, err := repo.existsSisProByCode(&models.Dci{}, data.Code)
	if err != nil {
		return err
	}

	result, isOk := resultData.(*models.Dci)
	if !isOk {
		log.Fatal("NO MAPPER TO Dci")
	}
	if result.Code == data.Code {
		return nil
	}
	return repo.db.Save(&data).Error
}
