package repositoryinterface

import "github.com/jdsierrab3991/scripts/34-golang-excel/domain/models"

type CupRipsRepositoryInterface interface {
	SaveCupRips(data *models.CupsRips) error
}
