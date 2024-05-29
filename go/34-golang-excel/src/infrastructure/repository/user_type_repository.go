package repository

import (
	"log"

	"github.com/jdsierrab3991/scripts/34-golang-excel/domain/models"
)

func (repo *Repository) SaveUserType(data *models.UserType) error {
	resultData, err := repo.existsSisProByCode(&models.UserType{}, data.Code)
	if err != nil {
		return err
	}

	result, isOk := resultData.(*models.UserType)
	if !isOk {
		log.Fatal("NO MAPPER TO UserType")
	}
	if result.Code == data.Code {
		return nil
	}
	return repo.db.Save(&data).Error
}
