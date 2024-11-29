package service

import (
	errorsrips "github.com/jdsierrab3991/scripts/34-golang-excel/domain/errors_rips"
	"github.com/jdsierrab3991/scripts/34-golang-excel/domain/libs"
	"github.com/jdsierrab3991/scripts/34-golang-excel/domain/models"
	repositoryinterface "github.com/jdsierrab3991/scripts/34-golang-excel/domain/repository_interface"
)

type GroupServiceService struct {
	repo repositoryinterface.GroupServiceRepositoryInterface
}

func NewGroupServiceService(repo repositoryinterface.GroupServiceRepositoryInterface) *GroupServiceService {
	return &GroupServiceService{
		repo: repo,
	}
}

func (GroupServiceService) GetCode() string {
	return libs.FFM
}

func (service *GroupServiceService) SaveSisproData(data interface{}) error {
	dataMapper, isOk := data.(*models.GroupService)
	if !isOk {
		return errorsrips.MapperError{Data: service.GetCode()}
	}
	return service.repo.SaveGroupService(dataMapper)
}

func (s *GroupServiceService) GetCodesForData() ([]string, error) {
	return s.repo.GetCodesForDataGroupService()
}
