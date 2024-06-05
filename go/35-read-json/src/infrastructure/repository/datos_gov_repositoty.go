package repository

import (
	"log"

	"github.com/jdsierrab3991/scripts/35-read-json/src/domain/model"
)

func (repo *Repository) SaveSubRegion(code, name, region string) error {
	var exists model.SubRegion
	err := repo.db.Where("code = ?", code).Find(&exists).Error
	if err != nil {
		return err
	}

	if exists.Code != "" {
		log.Printf("exisrst subregion %s", code)
		return nil
	}
	return repo.db.Save(&model.SubRegion{Name: name, Code: code, Region: region}).Error
}

func (repo *Repository) SaveMunicipality(code, name, codeSubregion string) error {
	var exists model.Municipality
	err := repo.db.Where("code = ?", code).Find(&exists).Error
	if err != nil {
		return err
	}

	if exists.Code != "" {
		log.Printf("exisrst municiaplity %s", code)
		return nil
	}
	return repo.db.Save(&model.Municipality{Name: name, Code: code, SubRegion: codeSubregion}).Error
}
