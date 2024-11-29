package repository

import (
	"log"

	"github.com/jdsierrab3991/scripts/34-golang-excel/domain/models"
)

func (repo *Repository) SaveCumSisPro(data *models.CumSispro) error {
	var result models.CumSispro
	err := repo.existsSisProByCodeGe(&result, data.Code)
	if err != nil {
		return err
	}

	if data.Code == result.Code {
		return nil
	}
	return repo.db.Save(&data).Error
}

func (repo *Repository) SaveCollectionConceptSisPro(data *models.CollectionConcept) error {
	resultData, err := repo.existsSisProByCode(&models.CollectionConcept{}, data.Code)
	if err != nil {
		return err
	}

	result, isOk := resultData.(*models.CollectionConcept)
	if !isOk {
		log.Fatal("NO MAPPER TO CollectionConcept")
	}
	if data.Code == result.Code {
		return nil
	}
	return repo.db.Save(&data).Error
}

func (repo *Repository) GetCodesForDataSispro() ([]string, error) {
	var result []string
	err := repo.GetCodesForData(&models.CumSispro{}, &result)
	if err != nil {
		return nil, err
	}
	return result, err
}

func (repo *Repository) GetCodesForCollectionConcept() ([]string, error) {
	var result []string
	err := repo.GetCodesForData(&models.CollectionConcept{}, &result)
	if err != nil {
		return nil, err
	}
	return result, err
}
