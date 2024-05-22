package repository

import (
	"log"

	"github.com/jdsierrab3991/scripts/34-golang-excel/domain/models"
)

func (repo *Repository) SaveCumSisPro(data *models.CumSispro) error {
	resultData, err := repo.existsSisProByCode(&models.CumSispro{}, data.Code)
	if err != nil {
		return err
	}

	result, isOk := resultData.(*models.CumSispro)
	if !isOk {
		log.Fatal("NO MAPPER TO GroupService")
	}
	if data.Code == result.Code {
		return nil
	}
	return repo.db.Save(&data).Error
}
