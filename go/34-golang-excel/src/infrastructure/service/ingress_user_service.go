package service

import (
	errorsrips "github.com/jdsierrab3991/scripts/34-golang-excel/domain/errors_rips"
	"github.com/jdsierrab3991/scripts/34-golang-excel/domain/libs"
	"github.com/jdsierrab3991/scripts/34-golang-excel/domain/models"
	repositoryinterface "github.com/jdsierrab3991/scripts/34-golang-excel/domain/repository_interface"
)

type IngressUserService struct {
	repo repositoryinterface.IngressUserRepositoryInterface
}

func NewIngressUserService(repo repositoryinterface.IngressUserRepositoryInterface) *IngressUserService {
	return &IngressUserService{
		repo: repo,
	}
}

func (IngressUserService) GetCode() string {
	return libs.FFM
}

func (service *IngressUserService) SaveSisproData(data interface{}) error {
	dataMapper, isOk := data.(*models.IngressUser)
	if !isOk {
		return errorsrips.MapperError{Data: service.GetCode()}
	}
	return service.repo.SaveIngressUser(dataMapper)
}
