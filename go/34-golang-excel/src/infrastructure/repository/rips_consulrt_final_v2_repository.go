package repository

import (
	"log"

	"github.com/jdsierrab3991/scripts/34-golang-excel/domain/models"
)

func (repo *Repository) SaveRipsConsultFinalV2(data *models.RipsConsultFinalV2) error {
	resultData, err := repo.existsSisProByCode(&models.RipsConsultFinalV2{}, data.Code)
	if err != nil {
		return err
	}

	result, isOk := resultData.(*models.RipsConsultFinalV2)
	if !isOk {
		log.Fatal("NO MAPPER TO RipsConsultFinalV2")
	}
	if result.Code == data.Code {
		return nil
	}
	return repo.db.Save(&data).Error
}
