package repository

import "github.com/jdsierrab3991/scripts/34-golang-excel/domain/models"

func (repo *Repository) SaveScrapp(code string) error {
	return repo.db.Save(&models.Scrapp{Type: "SIS_PROD", Code: code}).Error
}

func (repo *Repository) ExistsScrapp(code string) (*models.Scrapp, error) {
	var result *models.Scrapp
	err := repo.db.Model(&models.Scrapp{}).Where("code = ?", code).Find(&result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}
