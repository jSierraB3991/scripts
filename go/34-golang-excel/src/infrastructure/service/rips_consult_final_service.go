package service

import (
	errorsrips "github.com/jdsierrab3991/scripts/34-golang-excel/domain/errors_rips"
	"github.com/jdsierrab3991/scripts/34-golang-excel/domain/libs"
	"github.com/jdsierrab3991/scripts/34-golang-excel/domain/models"
	repositoryinterface "github.com/jdsierrab3991/scripts/34-golang-excel/domain/repository_interface"
)

type RipsConsultFinalService struct {
	repo repositoryinterface.RipsConsultFinalRepositoryInterface
}

func NewRipsConsultFinalService(repo repositoryinterface.RipsConsultFinalRepositoryInterface) *RipsConsultFinalService {
	return &RipsConsultFinalService{
		repo: repo,
	}
}

func (RipsConsultFinalService) GetCode() string {
	return libs.RIPSFinalidadConsultaVersion2
}

func (service *RipsConsultFinalService) SaveSisproData(data interface{}) error {
	dataMapper, isOk := data.(*models.RipsConsultFinalV2)
	if !isOk {
		return errorsrips.MapperError{Data: service.GetCode()}
	}
	return service.repo.SaveRipsConsultFinalV2(dataMapper)
}
