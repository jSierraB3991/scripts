package service

import (
	errorsrips "github.com/jdsierrab3991/scripts/34-golang-excel/domain/errors_rips"
	"github.com/jdsierrab3991/scripts/34-golang-excel/domain/libs"
	"github.com/jdsierrab3991/scripts/34-golang-excel/domain/models"
	repositoryinterface "github.com/jdsierrab3991/scripts/34-golang-excel/domain/repository_interface"
)

type CupRipsService struct {
	repo repositoryinterface.CupRipsRepositoryInterface
}

func NewCupRipsService(repo repositoryinterface.CupRipsRepositoryInterface) *CupRipsService {
	return &CupRipsService{
		repo: repo,
	}
}

func (CupRipsService) GetCode() string {
	return libs.CUPSRips
}

func (service *CupRipsService) SaveSisproData(data interface{}) error {
	dataMapper, isOk := data.(*models.CupsRips)
	if !isOk {
		return errorsrips.MapperError{Data: service.GetCode()}
	}
	return service.repo.SaveCupRips(dataMapper)
}

func (s *CupRipsService) GetCodesForData() ([]string, error) {
	return s.repo.GetCodesForCupsRips()
}
