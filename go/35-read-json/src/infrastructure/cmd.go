package infrastructure

import (
	"encoding/json"
	"io"
	"log"
	"os"

	"github.com/jdsierrab3991/scripts/35-read-json/src/domain/database"
	"github.com/jdsierrab3991/scripts/35-read-json/src/infrastructure/repository"
	"github.com/jdsierrab3991/scripts/35-read-json/src/infrastructure/rest"
	"github.com/jdsierrab3991/scripts/35-read-json/src/infrastructure/service"
	"gorm.io/gorm"
)

func Run(db *gorm.DB) {
	datas, err := readFile(os.Getenv("JSON_PATH"))
	if err != nil {
		log.Fatal(err)
	}
	repository := repository.InitiateRepo(db)
	database.AutoMigrate(repository)
	service := service.NewDatosGovService(repository)
	err = service.SaveAll(datas)
	if err != nil {
		log.Fatal(err)
	}
}

func readFile(jsonPath string) ([]rest.DatosGov, error) {
	jsonfile, err := os.Open(jsonPath)
	if err != nil {
		return nil, err
	}
	log.Println("Success open json")
	defer jsonfile.Close()

	bytevalue, _ := io.ReadAll(jsonfile)

	var result []rest.DatosGov
	err = json.Unmarshal(bytevalue, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
