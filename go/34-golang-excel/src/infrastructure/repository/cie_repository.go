package repository

import (
	"log"

	"github.com/jdsierrab3991/scripts/34-golang-excel/domain/models"
)

func (repo *Repository) SaveCie2036SisPro(data *models.Cie2036) error {
	resultData, err := repo.existsSisProByCode(&models.Cie2036{}, data.Code)
	if err != nil {
		return err
	}
	result, isOk := resultData.(*models.Cie2036)
	if !isOk {
		log.Fatal("NO MAPPER TO Cie2036")
	}
	if result.Code == data.Code {
		return nil
	}
	return repo.db.Save(&data).Error
}

func (repo *Repository) SaveCieSisPro(data *models.Cie) error {
	resultData, err := repo.existsSisProByCode(&models.Cie{}, data.Code)
	if err != nil {
		return err
	}
	result, isOk := resultData.(*models.Cie)
	if !isOk {
		log.Fatal("NO MAPPER TO Cie")
	}
	if result.Code == data.Code {
		return nil
	}
	return repo.db.Save(&data).Error
}
