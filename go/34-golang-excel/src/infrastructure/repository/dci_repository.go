package repository

import "github.com/jdsierrab3991/scripts/34-golang-excel/domain/models"

func (repo *Repository) SaveDci(data *models.Dci) error {
	result, err := repo.existsDci(data.Code)
	if err != nil {
		return err
	}
	if result.Code == data.Code {
		return nil
	}
	return repo.db.Save(&data).Error
}

func (repo *Repository) existsDci(code string) (*models.Dci, error) {
	var result *models.Dci
	err := repo.db.Model(&models.Dci{}).Where("code = ?", code).Find(&result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}
