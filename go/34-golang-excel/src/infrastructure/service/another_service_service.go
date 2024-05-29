package service

import (
	errorsrips "github.com/jdsierrab3991/scripts/34-golang-excel/domain/errors_rips"
	"github.com/jdsierrab3991/scripts/34-golang-excel/domain/libs"
	"github.com/jdsierrab3991/scripts/34-golang-excel/domain/models"
	repositoryinterface "github.com/jdsierrab3991/scripts/34-golang-excel/domain/repository_interface"
)

type AnotherServiceService struct {
	repository repositoryinterface.OtherServiceRepositoryInterface
}

func NewAnotherServiceService(repository repositoryinterface.OtherServiceRepositoryInterface) *AnotherServiceService {
	return &AnotherServiceService{
		repository: repository,
	}
}

func (AnotherServiceService) GetCode() string {
	return libs.TipoOtrosServicios
}

func (service *AnotherServiceService) SaveSisproData(data interface{}) error {
	dataMapper, isOk := data.(*models.OtherService)
	if !isOk {
		return errorsrips.MapperError{Data: service.GetCode()}
	}
	return service.repository.SaveOtherService(dataMapper)
}
