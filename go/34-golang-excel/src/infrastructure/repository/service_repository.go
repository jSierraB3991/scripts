package repository

import (
	"log"

	"github.com/jdsierrab3991/scripts/34-golang-excel/domain/models"
)

func (repo *Repository) SaveService(data *models.Service) error {
	resultData, err := repo.existsSisProByCode(&models.Service{}, data.Code)
	if err != nil {
		return err
	}

	result, isOk := resultData.(*models.Service)
	if !isOk {
		log.Fatal("NO MAPPER TO Service")
	}
	if result.Code == data.Code {
		return nil
	}
	return repo.db.Save(&data).Error
}
