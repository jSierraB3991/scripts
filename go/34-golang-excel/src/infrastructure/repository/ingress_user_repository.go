package repository

import (
	"log"

	"github.com/jdsierrab3991/scripts/34-golang-excel/domain/models"
)

func (repo *Repository) SaveIngressUser(data *models.IngressUser) error {
	resultData, err := repo.existsSisProByCode(&models.IngressUser{}, data.Code)
	if err != nil {
		return err
	}

	result, isOk := resultData.(*models.IngressUser)
	if !isOk {
		log.Fatal("NO MAPPER TO IngressUser")
	}
	if result.Code == data.Code {
		return nil
	}
	return repo.db.Save(&data).Error
}

func (repo *Repository) GetCodesForDataIngresUser() ([]string, error) {
	var result []string
	err := repo.GetCodesForData(&models.IngressUser{}, &result)
	if err != nil {
		return nil, err
	}
	return result, err
}
