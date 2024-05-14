package repository

import "github.com/jdsierrab3991/scripts/34-golang-excel/domain/models"

func (repo Repository) SaveCie2036SisPro(data *models.Cie2036) error {
	result, err := repo.existsCie2036(data.Code)
	if err != nil {
		return err
	}
	if result.Code == data.Code {
		return nil
	}
	return repo.db.Save(&data).Error
}
func (repo Repository) SaveCieSisPro(data *models.Cie) error {
	result, err := repo.existsCie(data.Code)
	if err != nil {
		return err
	}
	if result.Code == data.Code {
		return nil
	}
	return repo.db.Save(&data).Error
}

func (repo *Repository) existsCie2036(code string) (*models.Cie2036, error) {
	var result *models.Cie2036
	err := repo.db.Model(&models.Cie2036{}).Where("code = ?", code).Find(&result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (repo *Repository) existsCie(code string) (*models.Cie, error) {
	var result *models.Cie
	err := repo.db.Model(&models.Cie{}).Where("code = ?", code).Find(&result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}
