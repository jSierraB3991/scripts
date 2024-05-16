package repositoryinterface

import "github.com/jdsierrab3991/scripts/34-golang-excel/domain/models"

type DcirepositoryInterface interface {
	SaveDci(data *models.Dci) error
}
