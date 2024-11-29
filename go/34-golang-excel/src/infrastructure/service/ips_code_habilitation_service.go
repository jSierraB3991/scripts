package service

import (
	errorsrips "github.com/jdsierrab3991/scripts/34-golang-excel/domain/errors_rips"
	"github.com/jdsierrab3991/scripts/34-golang-excel/domain/libs"
	"github.com/jdsierrab3991/scripts/34-golang-excel/domain/models"
	repositoryinterface "github.com/jdsierrab3991/scripts/34-golang-excel/domain/repository_interface"
)

type IPSCodHabilitacionService struct {
	repo repositoryinterface.IPSCodHabilitacionRepositoryInterface
}

func NewIPSCodHabilitacionService(repo repositoryinterface.IPSCodHabilitacionRepositoryInterface) *IPSCodHabilitacionService {
	return &IPSCodHabilitacionService{
		repo: repo,
	}
}

func (IPSCodHabilitacionService) GetCode() string {
	return libs.FFM
}

func (service *IPSCodHabilitacionService) SaveSisproData(data interface{}) error {
	dataMapper, isOk := data.(*models.IpsCpdeHabilitation)
	if !isOk {
		return errorsrips.MapperError{Data: service.GetCode()}
	}
	return service.repo.SaveIPSCodHabilitacion(dataMapper)
}

func (s *IPSCodHabilitacionService) GetCodesForData() ([]string, error) {
	return s.repo.GetCodesForIpsCodeNoHabilitation()
}
