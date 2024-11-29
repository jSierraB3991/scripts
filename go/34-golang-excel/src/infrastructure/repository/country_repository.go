package repository

import (
	"log"

	"github.com/jdsierrab3991/scripts/34-golang-excel/domain/models"
)

func (repo *Repository) SaveCountry(data *models.Country) error {
	resultData, err := repo.existsSisProByCode(&models.Country{}, data.Code)
	if err != nil {
		return err
	}
	result, isOk := resultData.(*models.Country)
	if !isOk {
		log.Fatal("NO MAPPER TO Country")
	}
	if result.Code == data.Code {
		return nil
	}
	return repo.db.Save(&data).Error
}

func (repo *Repository) GetCodesForCountry() ([]string, error) {
	var result []string
	err := repo.GetCodesForData(&models.Country{}, &result)
	if err != nil {
		return nil, err
	}
	return result, err
}
