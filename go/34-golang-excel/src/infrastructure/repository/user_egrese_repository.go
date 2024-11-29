package repository

import (
	"log"

	"github.com/jdsierrab3991/scripts/34-golang-excel/domain/models"
)

func (repo *Repository) SaveUserEgrese(data *models.UserEgrese) error {
	resultData, err := repo.existsSisProByCode(&models.UserEgrese{}, data.Code)
	if err != nil {
		return err
	}

	result, isOk := resultData.(*models.UserEgrese)
	if !isOk {
		log.Fatal("NO MAPPER TO UserEgrese")
	}
	if data.Code == result.Code {
		return nil
	}
	return repo.db.Save(&data).Error
}

func (repo *Repository) GetCodesForUserEgrese() ([]string, error) {
	var result []string
	err := repo.GetCodesForData(&models.UserEgrese{}, &result)
	if err != nil {
		return nil, err
	}
	return result, err
}
