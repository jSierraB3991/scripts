package database

import (
	"encoding/base64"
	"log"
	"os"

	"cloud.google.com/go/cloudsqlconn"
	"cloud.google.com/go/cloudsqlconn/postgres/pgxv4"

	"gorm.io/gorm"

	_ "github.com/GoogleCloudPlatform/cloudsql-proxy/proxy/dialers/postgres"
	"github.com/jdsierrab3991/scripts/34-golang-excel/domain/models"
	"github.com/jdsierrab3991/scripts/34-golang-excel/infrastructure/repository"
	"gorm.io/driver/postgres"
)

func New(pg_url, gcpCredentialDbbase64 string) (*gorm.DB, func() error) {
	if len(gcpCredentialDbbase64) <= 0 {
		db, err := gorm.Open(postgres.Open(pg_url), &gorm.Config{})

		if err != nil {
			log.Println(err)
			os.Exit(1)
		}
		return db, nil
	}
	gcpCredentialDb, err := base64.StdEncoding.DecodeString(gcpCredentialDbbase64)
	if err != nil {
		log.Fatal(err)
	}

	cleanup, err := pgxv4.RegisterDriver("cloudsql-postgres", cloudsqlconn.WithCredentialsJSON(gcpCredentialDb))
	if err != nil {
		log.Fatal(err)
	}

	db, err := gorm.Open(postgres.New(postgres.Config{
		DriverName: "cloudsql-postgres",
		DSN:        pg_url,
	}))

	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	return db, cleanup

}

func AutoMigrate(repository *repository.Repository) {
	repository.GetDb().AutoMigrate(
		&models.Scrapp{},
		&models.AtentionModality{},
		&models.Cie{},
		&models.Cie2036{},
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
}
