package repository

import (
	"log"

	"github.com/jdsierrab3991/scripts/34-golang-excel/domain/models"
)

func (repo *Repository) SaveOtherService(data *models.OtherService) error {
	resultData, err := repo.existsSisProByCode(&models.OtherService{}, data.Code)
	if err != nil {
		return err
	}
	result, isOk := resultData.(*models.OtherService)
	if !isOk {
		log.Fatal("NO MAPPER TO OtherService")
	}
	if result.Code == data.Code {
		return nil
	}
	return repo.db.Save(&data).Error
}
