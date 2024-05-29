package repository

import (
	"log"

	"github.com/jdsierrab3991/scripts/34-golang-excel/domain/models"
)

func (repo *Repository) SaveMedicTypePOS(data *models.MedicTypePOS) error {
	resultData, err := repo.existsSisProByCode(&models.MedicTypePOS{}, data.Code)
	if err != nil {
		return err
	}

	result, isOk := resultData.(*models.MedicTypePOS)
	if !isOk {
		log.Fatal("NO MAPPER TO MedicTypePOS")
	}
	if result.Code == data.Code {
		return nil
	}
	return repo.db.Save(&data).Error
}
