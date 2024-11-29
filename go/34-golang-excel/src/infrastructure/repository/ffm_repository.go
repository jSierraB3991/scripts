package repository

import (
	"log"

	"github.com/jdsierrab3991/scripts/34-golang-excel/domain/models"
)

func (repo *Repository) SaveFfm(data *models.Ffm) error {
	resultData, err := repo.existsSisProByCode(&models.Ffm{}, data.Code)
	if err != nil {
		return err
	}

	result, isOk := resultData.(*models.Ffm)
	if !isOk {
		log.Fatal("NO MAPPER TO Ffm")
	}
	if result.Code == data.Code {
		return nil
	}
	return repo.db.Save(&data).Error
}

func (repo *Repository) GetCodesForFfm() ([]string, error) {
	var result []string
	err := repo.GetCodesForData(&models.Ffm{}, &result)
	if err != nil {
		return nil, err
	}
	return result, err
}
