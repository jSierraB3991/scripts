package repository

import "github.com/jdsierrab3991/scripts/34-golang-excel/domain/models"

func (repo *Repository) SaveUserEgrese(data *models.UserEgrese) error {
	result, err := repo.existsUserRegres(data.Code)
	if err != nil {
		return err
	}
	if data.Code == result.Code {
		return nil
	}
	return repo.db.Save(&data).Error
}

func (repo *Repository) existsUserRegres(code string) (*models.UserEgrese, error) {
	var result *models.UserEgrese
	err := repo.db.Model(&result).Where("code = ?", code).Find(&result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}
