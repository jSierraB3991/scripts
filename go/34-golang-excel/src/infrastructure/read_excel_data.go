package infrastructure

import (
	"fmt"
	"log"

	"github.com/jdsierrab3991/scripts/34-golang-excel/domain/libs"
	"github.com/jdsierrab3991/scripts/34-golang-excel/domain/mapper"
	serviceinterface "github.com/jdsierrab3991/scripts/34-golang-excel/domain/service_interface"
	"github.com/jdsierrab3991/scripts/34-golang-excel/infrastructure/repository"
	"github.com/jdsierrab3991/scripts/34-golang-excel/infrastructure/service"
	"github.com/xuri/excelize/v2"
)

type ReadExcelData struct {
	repo *repository.Repository
}

func NewReadExcelData(repo *repository.Repository) *ReadExcelData {
	return &ReadExcelData{
		repo: repo,
	}
}

func (read ReadExcelData) Run(homeFiles string, documents []string) error {

	for _, document := range documents {
		err := read.GetDataConfiguration(homeFiles, document)
		if err != nil {
			return err
		}
	}

	return nil
}

func (readDat *ReadExcelData) GetDataConfiguration(homeFiles, document string) error {

	fmt.Println(document)
	isValid := readDat.isForSave(document)
	if !isValid {
		return nil
	}

	f, err := excelize.OpenFile(homeFiles + document)
	if err != nil {
		return err
	}
	defer func() {
		//Close the spreadsheet.
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	//Get all the rows in the Sheet1.
	rows, err := f.GetRows("Table")
	if err != nil {
		return err
	}
	var code string
	for i, row := range rows {
		if i == 0 {
			for i := range row {
				var valueData string
				if len(rows[1]) > i {
					valueData = rows[1][i]
				}
				log.Printf("i: %v , name: %s, value: %s ", i, row[i], valueData)
			}
			continue
		} else if i == 1 {
			code = row[0]
			readDat.GetSisProService(code).SaveSisproData(mapper.GetDataSispro(row, code))
		} else {
			readDat.GetSisProService(code).SaveSisproData(mapper.GetDataSispro(row, code))
		}
	}
	return readDat.repo.SaveScrapp(document)
}

func (readData *ReadExcelData) isForSave(code string) bool {
	data, err := readData.repo.ExistsScrapp(code)
	if err != nil {
		return false
	}
	return data.Code != code
}

func (readDat ReadExcelData) GetSisProService(code string) serviceinterface.SisproServiceInterface {
	switch code {
	case libs.CatalogoCUMs:
		return service.NewCumSisProService(readDat.repo)
	case libs.CIE10Clasificacion2036:
		return service.NewCieService(readDat.repo)
	case libs.CIE10:
		return service.NewCieService(readDat.repo)
	case libs.CUPSRips:
		return service.NewCupRipsService(readDat.repo)
	case libs.CondicionyDestinoUsuarioEgreso:
		return service.NewUserEgreseService(readDat.repo)
	}
	return nil
}
