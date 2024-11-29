package repository

import (
	"log"

	"github.com/jdsierrab3991/scripts/34-golang-excel/domain/models"
)

func (repo *Repository) SaveGroupService(data *models.GroupService) error {
	resultData, err := repo.existsSisProByCode(&models.GroupService{}, data.Code)
	if err != nil {
		return err
	}

	result, isOk := resultData.(*models.GroupService)
	if !isOk {
		log.Fatal("NO MAPPER TO GroupService")
	}
	if result.Code == data.Code {
		return nil
	}
	return repo.db.Save(&data).Error
}

func (repo *Repository) GetCodesForDataGroupService() ([]string, error) {
	var result []string
	err := repo.GetCodesForData(&models.GroupService{}, &result)
	if err != nil {
		return nil, err
	}
	return result, err
}
