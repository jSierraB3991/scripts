package service

import (
	errorsrips "github.com/jdsierrab3991/scripts/34-golang-excel/domain/errors_rips"
	"github.com/jdsierrab3991/scripts/34-golang-excel/domain/libs"
	"github.com/jdsierrab3991/scripts/34-golang-excel/domain/models"
	repositoryinterface "github.com/jdsierrab3991/scripts/34-golang-excel/domain/repository_interface"
)

type DciService struct {
	repo repositoryinterface.DcirepositoryInterface
}

func NewDciService(repo repositoryinterface.DcirepositoryInterface) *DciService {
	return &DciService{
		repo: repo,
	}
}

func (DciService) GetCode() string {
	return libs.DCI
}

func (service *DciService) SaveSisproData(data interface{}) error {
	dataMapper, isOk := data.(*models.Dci)
	if !isOk {
		return errorsrips.MapperError{Data: service.GetCode()}
	}
	return service.repo.SaveDci(dataMapper)
}
