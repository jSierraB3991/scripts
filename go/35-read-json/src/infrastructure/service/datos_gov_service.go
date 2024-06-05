package service

import (
	"github.com/jdsierrab3991/scripts/35-read-json/src/infrastructure/repository"
	"github.com/jdsierrab3991/scripts/35-read-json/src/infrastructure/rest"
)

type DatosGovService struct {
	repo *repository.Repository
}

func NewDatosGovService(repo *repository.Repository) *DatosGovService {
	return &DatosGovService{repo: repo}
}

func (service *DatosGovService) save(data rest.DatosGov) error {
	err := service.repo.SaveSubRegion(data.CodeSubregion, data.SubRegion, data.Region)
	if err != nil {
		return err
	}

	return service.repo.SaveMunicipality(data.CodeMunicipality, data.Municipality, data.CodeSubregion)
}

func (service *DatosGovService) SaveAll(datas []rest.DatosGov) error {
	for _, data := range datas {
		err := service.save(data)
		if err != nil {
			return err
		}
	}
	return nil
}
