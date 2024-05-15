package repository

import "github.com/jdsierrab3991/scripts/34-golang-excel/domain/models"

func (repo *Repository) SaveCupRips(data *models.CupsRips) error {
	result, err := repo.existsCupRips(data.Code)
	if err != nil {
		return err
	}

	if result.Code == data.Code {
		return nil
	}
	return repo.db.Save(&data).Error
}

func (repo *Repository) existsCupRips(code string) (*models.CupsRips, error) {
	var result *models.CupsRips

	err := repo.db.Model(&models.CupsRips{}).Where("code = ?", code).Find(&result).Error
	if err != nil {
		return nil, err
	}

	return result, nil
}
