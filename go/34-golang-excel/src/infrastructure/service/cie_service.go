package service

import (
	errorsrips "github.com/jdsierrab3991/scripts/34-golang-excel/domain/errors_rips"
	"github.com/jdsierrab3991/scripts/34-golang-excel/domain/libs"
	"github.com/jdsierrab3991/scripts/34-golang-excel/domain/models"
	repositoryinterface "github.com/jdsierrab3991/scripts/34-golang-excel/domain/repository_interface"
)

type CieService struct {
	repository repositoryinterface.CieRepositoryInterface
}

func NewCieService(repository repositoryinterface.CieRepositoryInterface) *CieService {
	return &CieService{
		repository: repository,
	}
}

func (CieService) GetCode() string {
	return libs.CIE10Clasificacion2036
}

func (service *CieService) SaveSisproData(data interface{}) error {
	dataMapper, isOk := data.(*models.Cie)
	if !isOk {
		return service.SaveSisproData2036(data)
	}
	err := service.repository.SaveCieSisPro(dataMapper)
	if err != nil {
		return err
	}
	return nil
}

func (service *CieService) SaveSisproData2036(data interface{}) error {
	dataMapper, isOk := data.(*models.Cie2036)
	if !isOk {
		return errorsrips.MapperError{Data: service.GetCode()}
	}
	err := service.repository.SaveCie2036SisPro(dataMapper)
	if err != nil {
		return err
	}
	return nil
}

func (s *CieService) GetCodesForData() ([]string, error) {
	return s.repository.GetCodesForCie()
}
