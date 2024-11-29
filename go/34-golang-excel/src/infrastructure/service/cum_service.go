package service

import (
	errorsrips "github.com/jdsierrab3991/scripts/34-golang-excel/domain/errors_rips"
	"github.com/jdsierrab3991/scripts/34-golang-excel/domain/libs"
	"github.com/jdsierrab3991/scripts/34-golang-excel/domain/models"
	repositoryinterface "github.com/jdsierrab3991/scripts/34-golang-excel/domain/repository_interface"
)

type CumSisProService struct {
	repo repositoryinterface.CumRepositoryInterface
}

func (CumSisProService) GetCode() string {
	return libs.CatalogoCUMs
}

func NewCumSisProService(repo repositoryinterface.CumRepositoryInterface) *CumSisProService {
	return &CumSisProService{
		repo: repo,
	}
}

func (service *CumSisProService) SaveSisproData(data interface{}) error {
	dataMapper, isOk := data.(*models.CumSispro)
	if !isOk {
		return errorsrips.MapperError{Data: service.GetCode()}
	}
	return service.repo.SaveCumSisPro(dataMapper)
}

func (s *CumSisProService) GetCodesForData() ([]string, error) {
	return s.repo.GetCodesForDataSispro()
}
