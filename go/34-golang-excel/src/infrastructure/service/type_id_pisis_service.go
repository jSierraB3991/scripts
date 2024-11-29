package service

import (
	errorsrips "github.com/jdsierrab3991/scripts/34-golang-excel/domain/errors_rips"
	"github.com/jdsierrab3991/scripts/34-golang-excel/domain/libs"
	"github.com/jdsierrab3991/scripts/34-golang-excel/domain/models"
	repositoryinterface "github.com/jdsierrab3991/scripts/34-golang-excel/domain/repository_interface"
)

type TypeIdPISISService struct {
	repo repositoryinterface.TypeIdPISISRepositoryInterface
}

func NewTypeIdPISISService(repo repositoryinterface.TypeIdPISISRepositoryInterface) *TypeIdPISISService {
	return &TypeIdPISISService{
		repo: repo,
	}
}

func (TypeIdPISISService) GetCode() string {
	return libs.TipoIdPISIS
}

func (service *TypeIdPISISService) SaveSisproData(data interface{}) error {
	dataMapper, isOk := data.(*models.TypeIdPISIS)
	if !isOk {
		return errorsrips.MapperError{Data: service.GetCode()}
	}
	return service.repo.SaveTypeIdPISIS(dataMapper)
}

func (s *TypeIdPISISService) GetCodesForData() ([]string, error) {
	return s.repo.GetCodesForTypeIdPisis()
}
