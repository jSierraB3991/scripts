package repository

import (
	"log"

	"github.com/jdsierrab3991/scripts/34-golang-excel/domain/models"
)

func (repo *Repository) SaveIpsNoReps(data *models.IpsNoReps) error {
	resultData, err := repo.existsSisProByCode(&models.IpsNoReps{}, data.Code)
	if err != nil {
		return err
	}

	result, isOk := resultData.(*models.IpsNoReps)
	if !isOk {
		log.Fatal("NO MAPPER TO IpsNoReps")
	}
	if result.Code == data.Code {
		return nil
	}
	return repo.db.Save(&data).Error
}

func (repo *Repository) GetCodesForIpsNoReps() ([]string, error) {
	var result []string
	err := repo.GetCodesForData(&models.IpsNoReps{}, &result)
	if err != nil {
		return nil, err
	}
	return result, err
}
