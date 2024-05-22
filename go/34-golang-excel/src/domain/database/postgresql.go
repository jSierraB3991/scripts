package database

import (
	"log"
	"os"

	"github.com/jdsierrab3991/scripts/34-golang-excel/domain/models"
	"github.com/jdsierrab3991/scripts/34-golang-excel/infrastructure/repository"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func New(pg_url string) *gorm.DB {

	db, err := gorm.Open(postgres.Open(pg_url), &gorm.Config{})

	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	return db
}

func AutoMigrate(repository *repository.Repository) {
	repository.GetDb().AutoMigrate(
		&models.AtentionModality{},
		&models.Cie{},
		&models.Cie2036{},
		&models.CumSispro{},
		&models.CupsRips{},
		&models.Dci{},
		&models.Ffm{},
		&models.Ium{},
		&models.GroupService{},
		&models.IpsCpdeHabilitation{},
		&models.IpsNoReps{},
		&models.Municipality{},
		&models.Scrapp{},
		&models.UserEgrese{},
	)
}
