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
		log.Fatal("NO MAPPER TO Ium")
	}
	if result.Code == data.Code {
		return nil
	}
	return repo.db.Save(&data).Error
}

func (repo *Repository) GetCodesForIum() ([]string, error) {
	var result []string
	err := repo.GetCodesForData(&models.Ium{}, &result)
	if err != nil {
		return nil, err
	}
	return result, err
}
