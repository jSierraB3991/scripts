package repositoryinterface

import "github.com/jdsierrab3991/scripts/34-golang-excel/domain/models"

type CieRepositoryInterface interface {
	SaveCie2036SisPro(data *models.Cie2036) error
	SaveCieSisPro(data *models.Cie) error
}
