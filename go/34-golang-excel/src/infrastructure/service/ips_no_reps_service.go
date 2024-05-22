package service

import (
	errorsrips "github.com/jdsierrab3991/scripts/34-golang-excel/domain/errors_rips"
	"github.com/jdsierrab3991/scripts/34-golang-excel/domain/libs"
	"github.com/jdsierrab3991/scripts/34-golang-excel/domain/models"
	repositoryinterface "github.com/jdsierrab3991/scripts/34-golang-excel/domain/repository_interface"
)

type IpsNoReps struct {
	repo repositoryinterface.IpsNoRepsRepositoryInterface
}

func NewIpsNoReps(repo repositoryinterface.IpsNoRepsRepositoryInterface) *IpsNoReps {
	return &IpsNoReps{
		repo: repo,
	}
}

func (IpsNoReps) GetCode() string {
	return libs.FFM
}

func (service *IpsNoReps) SaveSisproData(data interface{}) error {
	dataMapper, isOk := data.(*models.IpsNoReps)
	if !isOk {
		return errorsrips.MapperError{Data: service.GetCode()}
	}
	return service.repo.SaveIpsNoReps(dataMapper)
}
