package repository

import (
	"log"

	"github.com/jdsierrab3991/scripts/34-golang-excel/domain/models"
)

func (repo *Repository) SaveRipsExternCauseV2(data *models.RipsCausaExternaV2) error {
	resultData, err := repo.existsSisProByCode(&models.RipsCausaExternaV2{}, data.Code)
	if err != nil {
		return err
	}

	result, isOk := resultData.(*models.RipsCausaExternaV2)
	if !isOk {
		log.Fatal("NO MAPPER TO RipsCausaExternaV2")
	}
	if result.Code == data.Code {
		return nil
	}
	return repo.db.Save(&data).Error
}

func (repo *Repository) GetCodesForRipsExternCauseV2() ([]string, error) {
	var result []string
	err := repo.GetCodesForData(&models.RipsCausaExternaV2{}, &result)
	if err != nil {
		return nil, err
	}
	return result, err
}
