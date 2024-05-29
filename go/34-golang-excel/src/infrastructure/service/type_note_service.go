package service

import (
	errorsrips "github.com/jdsierrab3991/scripts/34-golang-excel/domain/errors_rips"
	"github.com/jdsierrab3991/scripts/34-golang-excel/domain/libs"
	"github.com/jdsierrab3991/scripts/34-golang-excel/domain/models"
	repositoryinterface "github.com/jdsierrab3991/scripts/34-golang-excel/domain/repository_interface"
)

type TypeNoteService struct {
	repo repositoryinterface.TypeNotePOSRepositoryInterface
}

func NewTypeNoteService(repo repositoryinterface.TypeNotePOSRepositoryInterface) *TypeNoteService {
	return &TypeNoteService{
		repo: repo,
	}
}

func (TypeNoteService) GetCode() string {
	return libs.TipoNota
}

func (service *TypeNoteService) SaveSisproData(data interface{}) error {
	dataMapper, isOk := data.(*models.TypeNote)
	if !isOk {
		return errorsrips.MapperError{Data: service.GetCode()}
	}
	return service.repo.SaveTypeNote(dataMapper)
}
