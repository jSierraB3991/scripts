package service

import (
	errorsrips "github.com/jdsierrab3991/scripts/34-golang-excel/domain/errors_rips"
	"github.com/jdsierrab3991/scripts/34-golang-excel/domain/libs"
	"github.com/jdsierrab3991/scripts/34-golang-excel/domain/models"
	repositoryinterface "github.com/jdsierrab3991/scripts/34-golang-excel/domain/repository_interface"
)

type UprService struct {
	repo repositoryinterface.UPRRepositoryInterface
}

func NewUprService(repo repositoryinterface.UPRRepositoryInterface) *UprService {
	return &UprService{
		repo: repo,
	}
}

func (UprService) GetCode() string {
	return libs.UPR
}

func (service *UprService) SaveSisproData(data interface{}) error {
	dataMapper, isOk := data.(*models.UPR)
	if !isOk {
		return errorsrips.MapperError{Data: service.GetCode()}
	}
	return service.repo.SaveUPR(dataMapper)
}
