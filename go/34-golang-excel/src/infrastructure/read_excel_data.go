package infrastructure

import (
	"fmt"
	"log"

	"github.com/jdsierrab3991/scripts/34-golang-excel/domain/libs"
	"github.com/jdsierrab3991/scripts/34-golang-excel/domain/mapper"
	serviceinterface "github.com/jdsierrab3991/scripts/34-golang-excel/domain/service_interface"
	"github.com/jdsierrab3991/scripts/34-golang-excel/infrastructure/repository"
	"github.com/jdsierrab3991/scripts/34-golang-excel/infrastructure/service"
	"github.com/xuri/excelize/v2"
)

type ReadExcelData struct {
	repo *repository.Repository
}

func NewReadExcelData(repo *repository.Repository) *ReadExcelData {
	return &ReadExcelData{
		repo: repo,
	}
}

func (read ReadExcelData) Run(homeFiles string, documents []string) error {

	for _, document := range documents {
		err := read.GetDataConfiguration(homeFiles, document)
		if err != nil {
			return err
		}
	}

	return nil
}

func (readDat *ReadExcelData) GetDataConfiguration(homeFiles, document string) error {

	fmt.Println(document)
	isValid := readDat.isForSave(document)
	if !isValid {
		return nil
	}

	f, err := excelize.OpenFile(homeFiles + document)
	if err != nil {
		return err
	}
	defer func() {
		//Close the spreadsheet.
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	//Get all the rows in the Sheet1.
	rows, err := f.GetRows("Table")
	if err != nil {
		return err
	}
	var code string
	var service serviceinterface.SisproServiceInterface

	for i, row := range rows {
		if i == 0 {
			continue
		} else if i == 1 {
			code = row[0]
			service = readDat.GetSisProService(code)
		}

		if service == nil {
			log.Fatalf("el documento %s", document)
		}
		readDat.saveData(code, row, service)
	}
	log.Printf("SAVE DOCUMENT %s", document)
	return readDat.repo.SaveScrapp(document)
}

func (readDat *ReadExcelData) saveData(code string, row []string, service serviceinterface.SisproServiceInterface) {

	data := mapper.GetDataSispro(row, code)
	err := service.SaveSisproData(data)
	if err != nil {
		log.Fatal(err)
	}
}

func (readData *ReadExcelData) isForSave(code string) bool {
	data, err := readData.repo.ExistsScrapp(code)
	if err != nil {
		return false
	}
	return data.Code != code
}

func (readDat ReadExcelData) GetSisProService(code string) serviceinterface.SisproServiceInterface {
	switch code {
	case libs.CatalogoCUMs:
		return service.NewCumSisProService(readDat.repo)
	case libs.CIE10Clasificacion2036:
		return service.NewCieService(readDat.repo)
	case libs.CIE10:
		return service.NewCieService(readDat.repo)
	case libs.CUPSRips:
		return service.NewCupRipsService(readDat.repo)
	case libs.CondicionyDestinoUsuarioEgreso:
		return service.NewUserEgreseService(readDat.repo)
	case libs.DCI:
		return service.NewDciService(readDat.repo)
	case libs.FFM:
		return service.NewFfmService(readDat.repo)
	case libs.GrupoServicios:
		return service.NewGroupServiceService(readDat.repo)
	case libs.IPSCodHabilitacion:
		return service.NewIPSCodHabilitacionService(readDat.repo)
	case libs.IPSnoREPS:
		return service.NewIpsNoReps(readDat.repo)
	case libs.IUM:
		return service.NewIumService(readDat.repo)
	case libs.ModalidadAtencion:
		return service.NewModalityAtentionService(readDat.repo)
	case libs.RIPSCausaExternaVersion2:
		return service.NewRipsExternCauseV2Service(readDat.repo)
	case libs.RIPSFinalidadConsultaVersion2:
		return service.NewRipsConsultFinalService(readDat.repo)
	case libs.RIPSTipoDiagnosticoPrincipalVersion2:
		return service.NewRipsDiagnostictypePrincipalv2Service(readDat.repo)

	case libs.RIPSTipoUsuarioVersion2:
		return service.NewUserTypeService(readDat.repo)
	case libs.Servicios:
		return service.NewServiceService(readDat.repo)
	case libs.TipoIdPISIS:
		return service.NewTypeIdPISISService(readDat.repo)
	case libs.TipoMedicamentoPOSVersion2:
		return service.NewMedicTypePOSService(readDat.repo)
	case libs.TipoNota:
		return service.NewTypeNoteService(readDat.repo)
	case libs.TipoOtrosServicios:
		return service.NewAnotherServiceService(readDat.repo)
	case libs.UMM:
		return service.NewUmmService(readDat.repo)
	case libs.UPR:
		return service.NewUprService(readDat.repo)
	case libs.ViaIngresoUsuario:
		return service.NewIngressUserService(readDat.repo)

	}
	return nil
}
