package service

import (
	errorsrips "github.com/jdsierrab3991/scripts/34-golang-excel/domain/errors_rips"
	"github.com/jdsierrab3991/scripts/34-golang-excel/domain/libs"
	"github.com/jdsierrab3991/scripts/34-golang-excel/domain/models"
	repositoryinterface "github.com/jdsierrab3991/scripts/34-golang-excel/domain/repository_interface"
)

type MunicipalityService struct {
	repo repositoryinterface.MunicipalityRepositoryInterface
}

func NewMunicipalityService(repo repositoryinterface.MunicipalityRepositoryInterface) *MunicipalityService {
	return &MunicipalityService{
		repo: repo,
	}
}

func (MunicipalityService) GetCode() string {
	return libs.Municipio
}

func (service *MunicipalityService) SaveSisproData(data interface{}) error {
	dataMapper, isOk := data.(*models.Municipality)
	if !isOk {
		return errorsrips.MapperError{Data: service.GetCode()}
	}
	return service.repo.SaveMunicipality(dataMapper)
}
