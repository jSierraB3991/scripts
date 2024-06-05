package database

import (
	"github.com/jdsierrab3991/scripts/35-read-json/src/domain/model"
	"github.com/jdsierrab3991/scripts/35-read-json/src/infrastructure/repository"
)

func AutoMigrate(repo *repository.Repository) error {
	return repo.GetDb().AutoMigrate(
		&model.SubRegion{},
		&model.Municipality{},
	)
}
