package repository

import (
	"log"

	"github.com/jdsierrab3991/scripts/34-golang-excel/domain/models"
)

func (repo *Repository) SaveCupRips(data *models.CupsRips) error {
	resultData, err := repo.existsSisProByCode(&models.CupsRips{}, data.Code)
	if err != nil {
		return err
	}

	result, isOk := resultData.(*models.CupsRips)
	if !isOk {
		log.Fatal("NO MAPPER TO CupsRips")
	}
	if result.Code == data.Code {
		return nil
	}
	return repo.db.Save(&data).Error
}

func (repo *Repository) GetCodesForCupsRips() ([]string, error) {
	var result []string
	err := repo.GetCodesForData(&models.CupsRips{}, &result)
	if err != nil {
		return nil, err
	}
	return result, err
}
