package repositoryinterface

import "github.com/jdsierrab3991/scripts/34-golang-excel/domain/models"

type UserEgreRepositoryInterface interface {
	SaveUserEgrese(data *models.UserEgrese) error
}
