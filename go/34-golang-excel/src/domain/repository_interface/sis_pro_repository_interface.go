package repositoryinterface

import "github.com/jdsierrab3991/scripts/34-golang-excel/domain/models"

type CieRepositoryInterface interface {
	SaveCie2036SisPro(data *models.Cie2036) error
	SaveCieSisPro(data *models.Cie) error
}

type CumRepositoryInterface interface {
	SaveCumSisPro(data *models.CumSispro) error
}

type CupRipsRepositoryInterface interface {
	SaveCupRips(data *models.CupsRips) error
}

type DcirepositoryInterface interface {
	SaveDci(data *models.Dci) error
}

type UserEgreRepositoryInterface interface {
	SaveUserEgrese(data *models.UserEgrese) error
}

type FfmRepositoryInterface interface {
	SaveFfm(data *models.Ffm) error
}

type GroupServiceRepositoryInterface interface {
	SaveGroupService(data *models.GroupService) error
}
