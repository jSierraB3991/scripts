package database

import (
	"log"
	"os"

	"gorm.io/gorm"

	_ "github.com/GoogleCloudPlatform/cloudsql-proxy/proxy/dialers/postgres"
	"github.com/jdsierrab3991/scripts/34-golang-excel/domain/models"
	"github.com/jdsierrab3991/scripts/34-golang-excel/infrastructure/repository"
	"gorm.io/driver/postgres"
)

func NewPostgreSqlConnection(pg_url string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(pg_url), &gorm.Config{})

	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	return db

}

func AutoMigrate(repository *repository.Repository) {
	err := repository.GetDb().AutoMigrate(
		&models.Scrapp{},
		&models.AtentionModality{},
		&models.Cie{},
		&models.Cie2036{},
		&models.CollectionConcept{},
		&models.Country{},
		&models.CumSispro{},
		&models.CupsRips{},
		&models.Dci{},
		&models.Ffm{},
		&models.UserType{},
		&models.UserEgrese{},
		&models.IngressUser{},
		&models.GroupService{},
		&models.IpsCpdeHabilitation{},
		&models.IpsNoReps{},
		&models.Ium{},
		&models.MedicTypePOS{},
		&models.OtherService{},
		&models.RipsCausaExternaV2{},
		&models.RipsConsultFinalV2{},
		&models.RipsDiagnostictypePrincipalv2{},
		&models.Service{},

		&models.UPR{},
		&models.UMM{},
		&models.TypeNote{},
		&models.TypeIdPISIS{},
	)
	if err != nil {
		log.Fatal(err)
	}
}
