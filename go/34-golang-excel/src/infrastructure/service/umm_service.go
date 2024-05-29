package service

import (
	errorsrips "github.com/jdsierrab3991/scripts/34-golang-excel/domain/errors_rips"
	"github.com/jdsierrab3991/scripts/34-golang-excel/domain/libs"
	"github.com/jdsierrab3991/scripts/34-golang-excel/domain/models"
	repositoryinterface "github.com/jdsierrab3991/scripts/34-golang-excel/domain/repository_interface"
)

type UmmService struct {
	repo repositoryinterface.UMMRepositoryInterface
}

func NewUmmService(repo repositoryinterface.UMMRepositoryInterface) *UmmService {
	return &UmmService{
		repo: repo,
	}
}

func (UmmService) GetCode() string {
	return libs.UMM
}

func (service *UmmService) SaveSisproData(data interface{}) error {
	dataMapper, isOk := data.(*models.UMM)
	if !isOk {
		return errorsrips.MapperError{Data: service.GetCode()}
	}
	return service.repo.SaveUMM(dataMapper)
}
