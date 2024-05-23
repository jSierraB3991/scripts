package service

import (
	errorsrips "github.com/jdsierrab3991/scripts/34-golang-excel/domain/errors_rips"
	"github.com/jdsierrab3991/scripts/34-golang-excel/domain/libs"
	"github.com/jdsierrab3991/scripts/34-golang-excel/domain/models"
	repositoryinterface "github.com/jdsierrab3991/scripts/34-golang-excel/domain/repository_interface"
)

type RipsExternCauseV2Service struct {
	repo repositoryinterface.RipsExternCauseRepositoryInterface
}

func NewRipsExternCauseV2Service(repo repositoryinterface.RipsExternCauseRepositoryInterface) *RipsExternCauseV2Service {
	return &RipsExternCauseV2Service{
		repo: repo,
	}
}

func (RipsExternCauseV2Service) GetCode() string {
	return libs.RIPSCausaExternaVersion2
}

func (service *RipsExternCauseV2Service) SaveSisproData(data interface{}) error {
	dataMapper, isOk := data.(*models.RipsCausaExternaV2)
	if !isOk {
		return errorsrips.MapperError{Data: service.GetCode()}
	}
	return service.repo.SaveRipsExternCauseV2(dataMapper)
}
