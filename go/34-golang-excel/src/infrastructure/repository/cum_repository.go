package repository

import "github.com/jdsierrab3991/scripts/34-golang-excel/domain/models"

func (repo *Repository) SaveCumSisPro(data *models.CumSispro) error {
	model, err := repo.existsCumSis(data.Code)
	if err != nil {
		return err
	}

	if data.Code == model.Code {
		return nil
	}
	return repo.db.Save(&data).Error
}

func (repo *Repository) existsCumSis(code string) (*models.CumSispro, error) {
	var result *models.CumSispro
	err := repo.db.Model(&models.CumSispro{}).Where("code = ?", code).Find(&result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}
