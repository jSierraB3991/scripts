package repository

import (
	"log"

	"github.com/jdsierrab3991/scripts/34-golang-excel/domain/models"
)

func (repo *Repository) SaveAtentionModality(data *models.AtentionModality) error {
	resultData, err := repo.existsSisProByCode(&models.AtentionModality{}, data.Code)
	if err != nil {
		return err
	}
	result, isOk := resultData.(*models.AtentionModality)
	if !isOk {
		log.Fatal("NO MAPPER TO AtentionModality")
	}
	if result.Code == data.Code {
		return nil
	}
	return repo.db.Save(&data).Error
}

func (repo *Repository) GetCodesForAtentionModality() ([]string, error) {
	var result []string
	err := repo.GetCodesForData(&models.AtentionModality{}, &result)
	if err != nil {
		return nil, err
	}
	return result, err
}
