package service

import (
	errorsrips "github.com/jdsierrab3991/scripts/34-golang-excel/domain/errors_rips"
	"github.com/jdsierrab3991/scripts/34-golang-excel/domain/libs"
	"github.com/jdsierrab3991/scripts/34-golang-excel/domain/models"
	repositoryinterface "github.com/jdsierrab3991/scripts/34-golang-excel/domain/repository_interface"
)

type MedicTypePOSService struct {
	repo repositoryinterface.MedicTypePOSRepositoryInterface
}

func NewMedicTypePOSService(repo repositoryinterface.MedicTypePOSRepositoryInterface) *MedicTypePOSService {
	return &MedicTypePOSService{
		repo: repo,
	}
}

func (MedicTypePOSService) GetCode() string {
	return libs.TipoMedicamentoPOSVersion2
}

func (service *MedicTypePOSService) SaveSisproData(data interface{}) error {
	dataMapper, isOk := data.(*models.MedicTypePOS)
	if !isOk {
		return errorsrips.MapperError{Data: service.GetCode()}
	}
	return service.repo.SaveMedicTypePOS(dataMapper)
}

func (s *MedicTypePOSService) GetCodesForData() ([]string, error) {
	return s.repo.GetCodesForMedicTypeOs()
}
