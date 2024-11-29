package repository

import (
	"log"

	"github.com/jdsierrab3991/scripts/34-golang-excel/domain/models"
)

func (repo *Repository) SaveTypeIdPISIS(data *models.TypeIdPISIS) error {
	resultData, err := repo.existsSisProByCode(&models.TypeIdPISIS{}, data.Code)
	if err != nil {
		return err
	}

	result, isOk := resultData.(*models.TypeIdPISIS)
	if !isOk {
		log.Fatal("NO MAPPER TO TypeIdPISIS")
	}
	if result.Code == data.Code {
		return nil
	}
	return repo.db.Save(&data).Error
}

func (repo *Repository) GetCodesForTypeIdPisis() ([]string, error) {
	var result []string
	err := repo.GetCodesForData(&models.TypeIdPISIS{}, &result)
	if err != nil {
		return nil, err
	}
	return result, err
}
