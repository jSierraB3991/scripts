package service

import (
	errorsrips "github.com/jdsierrab3991/scripts/34-golang-excel/domain/errors_rips"
	"github.com/jdsierrab3991/scripts/34-golang-excel/domain/libs"
	"github.com/jdsierrab3991/scripts/34-golang-excel/domain/models"
	repositoryinterface "github.com/jdsierrab3991/scripts/34-golang-excel/domain/repository_interface"
)

type IumService struct {
	repo repositoryinterface.IumsRepositoryInterface
}

func NewIumService(repo repositoryinterface.IumsRepositoryInterface) *IumService {
	return &IumService{
		repo: repo,
	}
}

func (IumService) GetCode() string {
	return libs.FFM
}

func (service *IumService) SaveSisproData(data interface{}) error {
	dataMapper, isOk := data.(*models.Ium)
	if !isOk {
		return errorsrips.MapperError{Data: service.GetCode()}
	}
	return service.repo.SaveIum(dataMapper)
}
func (s *IumService) GetCodesForData() ([]string, error) {
	return s.repo.GetCodesForIum()
}
