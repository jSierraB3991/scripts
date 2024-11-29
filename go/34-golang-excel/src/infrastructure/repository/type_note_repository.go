package repository

import (
	"log"

	"github.com/jdsierrab3991/scripts/34-golang-excel/domain/models"
)

func (repo *Repository) SaveTypeNote(data *models.TypeNote) error {

	resultData, err := repo.existsSisProByCode(&models.TypeNote{}, data.Code)
	if err != nil {
		return err
	}

	result, isOk := resultData.(*models.TypeNote)
	if !isOk {
		log.Fatal("NO MAPPER TO TypeNote")
	}
	if result.Code == data.Code {
		return nil
	}
	return repo.db.Save(&data).Error
}

func (repo *Repository) GetCodesForTypeNote() ([]string, error) {
	var result []string
	err := repo.GetCodesForData(&models.TypeNote{}, &result)
	if err != nil {
		return nil, err
	}
	return result, err
}
