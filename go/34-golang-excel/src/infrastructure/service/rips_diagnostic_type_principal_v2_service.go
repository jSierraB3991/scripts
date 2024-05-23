package service

import (
	errorsrips "github.com/jdsierrab3991/scripts/34-golang-excel/domain/errors_rips"
	"github.com/jdsierrab3991/scripts/34-golang-excel/domain/libs"
	"github.com/jdsierrab3991/scripts/34-golang-excel/domain/models"
	repositoryinterface "github.com/jdsierrab3991/scripts/34-golang-excel/domain/repository_interface"
)

type RipsDiagnostictypePrincipalv2Service struct {
	repo repositoryinterface.RipsDiagnosticTypePincipalRepositoryInterface
}

func NewRipsDiagnostictypePrincipalv2Service(repo repositoryinterface.RipsDiagnosticTypePincipalRepositoryInterface) *RipsDiagnostictypePrincipalv2Service {
	return &RipsDiagnostictypePrincipalv2Service{
		repo: repo,
	}
}

func (RipsDiagnostictypePrincipalv2Service) GetCode() string {
	return libs.RIPSTipoDiagnosticoPrincipalVersion2
}

func (service *RipsDiagnostictypePrincipalv2Service) SaveSisproData(data interface{}) error {
	dataMapper, isOk := data.(*models.RipsDiagnostictypePrincipalv2)
	if !isOk {
		return errorsrips.MapperError{Data: service.GetCode()}
	}
	return service.repo.SaveRipsDiagnostictypePrincipalv2(dataMapper)
}
