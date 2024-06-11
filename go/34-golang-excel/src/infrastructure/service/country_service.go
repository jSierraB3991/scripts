package service

import (
	errorsrips "github.com/jdsierrab3991/scripts/34-golang-excel/domain/errors_rips"
	"github.com/jdsierrab3991/scripts/34-golang-excel/domain/libs"
	"github.com/jdsierrab3991/scripts/34-golang-excel/domain/models"
	repositoryinterface "github.com/jdsierrab3991/scripts/34-golang-excel/domain/repository_interface"
)

type CountryService struct {
	repo repositoryinterface.CountryRepositoryInterface
}

func NewCountryService(repo repositoryinterface.CountryRepositoryInterface) *CountryService {
	return &CountryService{
		repo: repo,
	}
}

func (CountryService) GetCode() string {
	return libs.Pais
}

func (service *CountryService) SaveSisproData(data interface{}) error {
	dataMapper, isOk := data.(*models.Country)
	if !isOk {
		return errorsrips.MapperError{Data: service.GetCode()}
	}
	return service.repo.SaveCountry(dataMapper)
}
