package service

import (
	errorsrips "github.com/jdsierrab3991/scripts/34-golang-excel/domain/errors_rips"
	"github.com/jdsierrab3991/scripts/34-golang-excel/domain/libs"
	"github.com/jdsierrab3991/scripts/34-golang-excel/domain/models"
	repositoryinterface "github.com/jdsierrab3991/scripts/34-golang-excel/domain/repository_interface"
)

func (UserEgreseService) GetCode() string {
	return libs.CondicionyDestinoUsuarioEgreso
}

type UserEgreseService struct {
	repo repositoryinterface.UserEgreRepositoryInterface
}

func NewUserEgreseService(repo repositoryinterface.UserEgreRepositoryInterface) *UserEgreseService {
	return &UserEgreseService{
		repo: repo,
	}
}

func (service *UserEgreseService) SaveSisproData(data interface{}) error {
	dataMapper, isOk := data.(*models.UserEgrese)
	if !isOk {
		return errorsrips.MapperError{Data: service.GetCode()}
	}
	return service.repo.SaveUserEgrese(dataMapper)
}

func (s *UserEgreseService) GetCodesForData() ([]string, error) {
	return s.repo.GetCodesForUserEgrese()
}
