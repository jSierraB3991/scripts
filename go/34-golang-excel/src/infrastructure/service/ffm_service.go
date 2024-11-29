package service

import (
	errorsrips "github.com/jdsierrab3991/scripts/34-golang-excel/domain/errors_rips"
	"github.com/jdsierrab3991/scripts/34-golang-excel/domain/libs"
	"github.com/jdsierrab3991/scripts/34-golang-excel/domain/models"
	repositoryinterface "github.com/jdsierrab3991/scripts/34-golang-excel/domain/repository_interface"
)

type FfmService struct {
	repo repositoryinterface.FfmRepositoryInterface
}

func NewFfmService(repo repositoryinterface.FfmRepositoryInterface) *FfmService {
	return &FfmService{
		repo: repo,
	}
}

func (FfmService) GetCode() string {
	return libs.FFM
}

func (service *FfmService) SaveSisproData(data interface{}) error {
	dataMapper, isOk := data.(*models.Ffm)
	if !isOk {
		return errorsrips.MapperError{Data: service.GetCode()}
	}
	return service.repo.SaveFfm(dataMapper)
}

func (s *FfmService) GetCodesForData() ([]string, error) {
	return s.repo.GetCodesForFfm()
}
