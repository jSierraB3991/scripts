package service

import (
	errorsrips "github.com/jdsierrab3991/scripts/34-golang-excel/domain/errors_rips"
	"github.com/jdsierrab3991/scripts/34-golang-excel/domain/libs"
	"github.com/jdsierrab3991/scripts/34-golang-excel/domain/models"
	repositoryinterface "github.com/jdsierrab3991/scripts/34-golang-excel/domain/repository_interface"
)

type UserTypeService struct {
	repo repositoryinterface.UserTypeRepositoryInterface
}

func NewUserTypeService(repo repositoryinterface.UserTypeRepositoryInterface) *UserTypeService {
	return &UserTypeService{
		repo: repo,
	}
}

func (UserTypeService) GetCode() string {
	return libs.RIPSTipoUsuarioVersion2
}

func (service *UserTypeService) SaveSisproData(data interface{}) error {
	dataMapper, isOk := data.(*models.UserType)
	if !isOk {
		return errorsrips.MapperError{Data: service.GetCode()}
	}
	return service.repo.SaveUserType(dataMapper)
}
