package service

import (
	errorsrips "github.com/jdsierrab3991/scripts/34-golang-excel/domain/errors_rips"
	"github.com/jdsierrab3991/scripts/34-golang-excel/domain/libs"
	"github.com/jdsierrab3991/scripts/34-golang-excel/domain/models"
	repositoryinterface "github.com/jdsierrab3991/scripts/34-golang-excel/domain/repository_interface"
)

type ServiceService struct {
	repo repositoryinterface.ServiceRepositoryInterface
}

func NewServiceService(repo repositoryinterface.ServiceRepositoryInterface) *ServiceService {
	return &ServiceService{
		repo: repo,
	}
}

func (ServiceService) GetCode() string {
	return libs.Servicios
}

func (service *ServiceService) SaveSisproData(data interface{}) error {
	dataMapper, isOk := data.(*models.Service)
	if !isOk {
		return errorsrips.MapperError{Data: service.GetCode()}
	}
	return service.repo.SaveService(dataMapper)
}

func (s *ServiceService) GetCodesForData() ([]string, error) {
	return s.repo.GetCodesForService()
}
