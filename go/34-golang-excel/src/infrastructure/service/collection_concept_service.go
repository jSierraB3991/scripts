package service

import (
	errorsrips "github.com/jdsierrab3991/scripts/34-golang-excel/domain/errors_rips"
	"github.com/jdsierrab3991/scripts/34-golang-excel/domain/libs"
	"github.com/jdsierrab3991/scripts/34-golang-excel/domain/models"
	repositoryinterface "github.com/jdsierrab3991/scripts/34-golang-excel/domain/repository_interface"
)

type CollectionConceptService struct {
	repo repositoryinterface.CollectionConceptRepositoryInterface
}

func (CollectionConceptService) GetCode() string {
	return libs.ConceptoRecaudo
}

func NewCollectionConceptService(repo repositoryinterface.CollectionConceptRepositoryInterface) *CollectionConceptService {
	return &CollectionConceptService{
		repo: repo,
	}
}

func (service *CollectionConceptService) SaveSisproData(data interface{}) error {
	dataMapper, isOk := data.(*models.CollectionConcept)
	if !isOk {
		return errorsrips.MapperError{Data: service.GetCode()}
	}
	err := service.repo.SaveCollectionConceptSisPro(dataMapper)
	if err != nil {
		return err
	}
	return nil
}

func (s *CollectionConceptService) GetCodesForData() ([]string, error) {
	return s.repo.GetCodesForCollectionConcept()
}
