package repositoryinterface

import "github.com/jdsierrab3991/scripts/34-golang-excel/domain/models"

type CumRepositoryInterface interface {
	SaveCumSisPro(data *models.CumSispro) error
}
