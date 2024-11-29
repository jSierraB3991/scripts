package service

import (
	errorsrips "github.com/jdsierrab3991/scripts/34-golang-excel/domain/errors_rips"
	"github.com/jdsierrab3991/scripts/34-golang-excel/domain/libs"
	"github.com/jdsierrab3991/scripts/34-golang-excel/domain/models"
	repositoryinterface "github.com/jdsierrab3991/scripts/34-golang-excel/domain/repository_interface"
)

type ModalityAtentionService struct {
	repo repositoryinterface.ModalityAtentionRepositoryInterface
}

func NewModalityAtentionService(repo repositoryinterface.ModalityAtentionRepositoryInterface) *ModalityAtentionService {
	return &ModalityAtentionService{
		repo: repo,
	}
}

func (ModalityAtentionService) GetCode() string {
	return libs.ModalidadAtencion
}

func (service *ModalityAtentionService) SaveSisproData(data interface{}) error {
	dataMapper, isOk := data.(*models.AtentionModality)
	if !isOk {
		return errorsrips.MapperError{Data: service.GetCode()}
	}
	return service.repo.SaveAtentionModality(dataMapper)
}
func (s ModalityAtentionService) GetCodesForData() ([]string, error) {
	return s.repo.GetCodesForAtentionModality()
}
