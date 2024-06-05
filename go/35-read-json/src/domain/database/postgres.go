package database

import (
	"encoding/base64"
	"log"
	"os"

	"cloud.google.com/go/cloudsqlconn"
	"cloud.google.com/go/cloudsqlconn/postgres/pgxv4"

	"gorm.io/gorm"

	_ "github.com/GoogleCloudPlatform/cloudsql-proxy/proxy/dialers/postgres"
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
