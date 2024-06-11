package mapper

import (
	"log"
	"time"

	errorsrips "github.com/jdsierrab3991/scripts/34-golang-excel/domain/errors_rips"
	"github.com/jdsierrab3991/scripts/34-golang-excel/domain/libs"
	"github.com/jdsierrab3991/scripts/34-golang-excel/domain/models"
)

func GetDataSispro(row []string, code string) interface{} {
	switch code {
	case libs.CatalogoCUMs:
		return getDatalogoCums(row)
	case libs.CIE10:
		return getDataCie10(row)
	case libs.CIE10Clasificacion2036:
		return getDataCie102036(row)
	case libs.CUPSRips:
		return getDataCupRips(row)
	case libs.CondicionyDestinoUsuarioEgreso:
		return getDataUserEgrese(row)
	case libs.DCI:
		return getDataDci(row)
	case libs.FFM:
		return getDataFfm(row)
	case libs.GrupoServicios:
		return getGroupService(row)
	case libs.IPSCodHabilitacion:
		return getIpsCpdeHabilitation(row)
	case libs.IPSnoREPS:
		return getIpsNoReps(row)
	case libs.IUM:
		return getIum(row)
	case libs.ModalidadAtencion:
		return getAtentionModality(row)
	case libs.RIPSCausaExternaVersion2:
		return getRipsCausaExternaV2(row)
	case libs.RIPSFinalidadConsultaVersion2:
		return getRipsFinalidadConsultaV2(row)
	case libs.RIPSTipoDiagnosticoPrincipalVersion2:
		return getRipsDiagnostictypePrincipalv2(row)
	case libs.RIPSTipoUsuarioVersion2:
		return getRIPSTipoUsuarioVersion2(row)
	case libs.Servicios:
		return getService(row)
	case libs.TipoIdPISIS:
		return getTypeIdPisis(row)
	case libs.TipoMedicamentoPOSVersion2:
		return getMedicTypePOS(row)
	case libs.TipoNota:
		return getTypeNote(row)
	case libs.TipoOtrosServicios:
		return getAnotherService(row)
	case libs.UMM:
		return getUmm(row)
	case libs.UPR:
		return getUpr(row)
	case libs.ViaIngresoUsuario:
		return getIngressUser(row)
	case libs.Pais:
		return getCountry(row)
	}
	return nil
}

func updateSisproFormat(date string) *time.Time {
	updateFormat, err := time.Parse(libs.SISPRO_DATE_FORMAT, date)
	if err != nil {
		return nil
	}
	return &updateFormat
}

func isPublicPrivate(row []string, index int) *bool {
	if len(row) > index {
		result := row[index] == "False"
		return &result
	}
	return nil
}

func validateVoidData(rows []string, indexs []int) error {
	for _, i := range indexs {
		if len(rows) > i && rows[i] != "" && rows[i] != "NULL" {
			return errorsrips.ValidataInMapperSisPro{Data: rows[i]}
		}
	}
	return nil
}

func getPointerString(data string) *string {
	if data == "NULL" {
		return nil
	}
	return &data
}

func getDataUserEgrese(rows []string) *models.UserEgrese {
	err := validateVoidData(rows, []int{3, 5, 13, 14, 15, 16, 17, 18, 19, 21})
	if err != nil {
		log.Panic(err)
	}
	return &models.UserEgrese{
		Code:            rows[1],
		Name:            rows[2],
		IsAvailable:     rows[4] == "SI",
		IsStandartGel:   rows[6] == "True",
		IsStandardMSPS:  rows[7] == "True",
		Consulation:     rows[8],
		Procedure:       rows[9],
		Emergency:       rows[10] == "SI",
		Hospitalization: rows[11] == "SI",
		Born:            rows[12] == "SI",
		UpdateDate:      *updateSisproFormat(rows[20]),
	}
}

func getRIPSTipoUsuarioVersion2(row []string) *models.UserType {

	err := validateVoidData(row, []int{3, 5, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19})
	if err != nil {
		log.Panic(err)
	}

	return &models.UserType{
		Code:           row[1],
		Name:           row[2],
		IsAvailable:    row[4] == "SI",
		IsStandartGel:  row[6] == "True",
		IsStandardMSPS: row[7] == "True",
		UpdateDate:     *updateSisproFormat(row[20]),
	}
}

func getDataCie10(rows []string) *models.Cie {
	err := validateVoidData(rows, []int{5, 14, 18, 19})
	if err != nil {
		log.Panic(err)
	}

	return &models.Cie{
		Code:            rows[1],
		Name:            rows[2],
		Description:     rows[3],
		IsAvailable:     rows[4] == "SI",
		IsStandartGel:   rows[6] == "True",
		IsStandardMSPS:  rows[7] == "True",
		AppliesToSex:    libs.GetUintForString(rows[8]),
		MinimunAge:      libs.GetUintForString(rows[9]),
		MaximunAge:      libs.GetUintForString(rows[10]),
		MortalityGroup:  libs.GetUintForString(rows[11]),
		ExtraV:          rows[12],
		Chapter:         rows[13],
		Subgroup:        libs.GetUintForString(rows[15]),
		Category:        libs.GetUintForString(rows[16]),
		Sex:             rows[17],
		UpdateDate:      *updateSisproFormat(rows[20]),
		IsPublicPrivate: isPublicPrivate(rows, 21),
	}
}

func getDataCie102036(rows []string) *models.Cie2036 {
	err := validateVoidData(rows, []int{5, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19})
	if err != nil {
		log.Panic(err)
	}
	return &models.Cie2036{
		Code:            rows[1],
		Name:            rows[2],
		Description:     rows[3],
		IsAvailable:     rows[4] == "SI",
		IsStandartGel:   rows[6] == "True",
		IsStandardMSPS:  rows[7] == "True",
		UpdateDate:      *updateSisproFormat(rows[20]),
		IsPublicPrivate: isPublicPrivate(rows, 21),
	}
}

func getDatalogoCums(rows []string) *models.CumSispro {
	err := validateVoidData(rows, []int{5, 17, 18, 19, 21})
	if err != nil {
		log.Panic(err)
	}
	return &models.CumSispro{
		Code:                  rows[1],
		Name:                  rows[2],
		Description:           rows[3],
		Available:             rows[4] == "SI",
		IsStandartGel:         rows[6] == "True",
		IsStandardMSPS:        rows[7] == "True",
		IndicatorSampleMedic:  rows[8] == "SI",
		AtcCode:               rows[9],
		ATC:                   rows[10],
		HealthRegister:        rows[11],
		ActivePrinciple:       rows[12],
		AmountActivePrinciple: rows[13],
		ViaAdministratio:      rows[15],
		AmountPresentation:    libs.GetUintForString(rows[16]),
		UpdateDate:            *updateSisproFormat(rows[20]),
	}
}

func getDataCupRips(rows []string) *models.CupsRips {
	err := validateVoidData(rows, []int{5, 18, 19, 21})
	if err != nil {
		log.Panic(err)
	}

	return &models.CupsRips{
		Code:           rows[1],
		Name:           rows[2],
		Description:    rows[3],
		IsAvailable:    rows[4] == "SI",
		IsStandartGel:  rows[6] == "True",
		IsStandardMSPS: rows[7] == "True",
		CupCode:        rows[8],
		Qx:             rows[9],
		MinimunNumber:  libs.GetUintForString(rows[10]),
		MaximunNumber:  libs.GetUintForString(rows[11]),
		DxRequired:     rows[12],
		Sex:            getPointerString(rows[13]),
		Ambit:          getPointerString(rows[14]),
		Stay:           getPointerString(rows[15]),
		Coverage:       getPointerString(rows[16]),
		Duplicate:      getPointerString(rows[17]),
	}
}

func getDataDci(rows []string) *models.Dci {

	err := validateVoidData(rows, []int{3, 5, 9, 10, 11, 12, 14, 15, 16, 17, 18, 19, 21})
	if err != nil {
		log.Panic(err)
	}
	return &models.Dci{
		Code:           rows[1],
		Name:           rows[2],
		IsAvailable:    rows[4] == "SI",
		IsStandartGel:  rows[6] == "True",
		IsStandardMSPS: rows[7] == "True",
		Extra:          libs.GetStringPoint(rows[8]),
		ExtraVI:        libs.GetUintPoint(rows[13]),
		UpdateDate:     *updateSisproFormat(rows[20]),
	}
}

func getDataFfm(rows []string) *models.Ffm {
	err := validateVoidData(rows, []int{5, 8, 16, 17, 18, 19, 21})
	if err != nil {
		log.Panic(err)
	}
	return &models.Ffm{
		Code:                  rows[1],
		Name:                  rows[2],
		Description:           rows[3],
		IsAvailable:           rows[4] == "SI",
		IsStandartGel:         rows[6] == "True",
		IsStandardMSPS:        rows[7] == "True",
		Level2Group:           rows[9] == "VERDADERO",
		Level2GroupDefinition: rows[10],
		Level3Group:           rows[11],
		Level3GroupDefinition: rows[12],
		ExtraVI:               rows[13],
		ExtraVII:              rows[14],
		UpdateDate:            *updateSisproFormat(rows[20]),
	}
}

func getGroupService(rows []string) *models.GroupService {
	err := validateVoidData(rows, []int{3, 5, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 21})
	if err != nil {
		log.Panic(err)
	}
	return &models.GroupService{
		Code:           rows[1],
		Name:           rows[2],
		IsAvailable:    rows[4] == "SI",
		IsStandartGel:  rows[6] == "SI",
		IsStandardMSPS: rows[7] == "SI",
		UpdateDate:     *updateSisproFormat(rows[20]),
	}
}

func getIpsCpdeHabilitation(row []string) *models.IpsCpdeHabilitation {
	err := validateVoidData(row, []int{3, 5, 16, 17, 18, 19, 21})
	if err != nil {
		log.Panic(err)
	}
	return &models.IpsCpdeHabilitation{
		Code:           row[1],
		Name:           row[2],
		IsAvailable:    row[4] == "SI",
		IsStandartGel:  row[6] == "SI",
		IsStandardMSPS: row[7] == "SI",
		TypeIdPres:     row[8],
		NumIdPres:      row[9],
		CodePres:       row[10],
		CodeMpiSede:    row[11],
		NameMpiSede:    row[12],
		NameDeptoSede:  row[13],
		ClassPres:      libs.GetUintForString(row[14]),
		NameClassPres:  row[15],
		UpdateDate:     *updateSisproFormat(row[20]),
	}
}

func getIpsNoReps(row []string) *models.IpsNoReps {
	err := validateVoidData(row, []int{5, 18, 19})
	if err != nil {
		log.Panic(err)
	}
	return &models.IpsNoReps{
		Code:             row[1],
		Name:             row[2],
		Description:      row[3],
		IsAvailable:      row[4] == "SI",
		IsStandartGel:    row[6] == "Verdadero",
		IsStandardMSPS:   row[7] == "Verdadero",
		Telphone:         row[8],
		Manager:          row[9],
		Regime:           row[10],
		CodeDepto:        row[11],
		Department:       row[12],
		CodeMunicipality: row[13],
		Municipality:     row[14],
		IpsType:          row[15],
		AtentionLevel:    libs.GetUintForString(row[16]),
		Nit:              row[17],
		UpdateDate:       *updateSisproFormat(row[20]),
		IsPublicPrivate:  isPublicPrivate(row, 21),
	}
}

func getIum(row []string) *models.Ium {
	err := validateVoidData(row, []int{3, 5, 18, 19, 21})
	if err != nil {
		log.Panic(err)
	}
	return &models.Ium{
		Code:                         row[1],
		Name:                         row[2],
		IsAvailable:                  row[4] == "SI",
		IsStandartGel:                row[6] == "Verdadero",
		IsStandardMSPS:               row[7] == "Verdadero",
		Nivel1:                       row[8],
		ActivePrincipal:              row[9],
		CodeActivePrincipal:          row[10],
		PharmaceuticalForm:           row[11],
		CodePharmaceuticalForm:       row[12],
		Nivel2:                       libs.GetUintForString(row[13]),
		CodeComercialitionForm:       row[14],
		Nivel3:                       libs.GetUintForString(row[15]),
		ConditionResgiterMedicSample: row[16],
		PackegeUnique:                row[17],
		UpdateDate:                   *updateSisproFormat(row[20]),
	}
}

func getAtentionModality(row []string) *models.AtentionModality {

	err := validateVoidData(row, []int{3, 5, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 21})
	if err != nil {
		log.Panic(err)
	}
	return &models.AtentionModality{
		Code:           row[1],
		Name:           row[2],
		IsAvailable:    row[4] == "SI",
		IsStandartGel:  row[6] == "Verdadero",
		IsStandardMSPS: row[7] == "Verdadero",
		UpdateDate:     *updateSisproFormat(row[20]),
	}
}
func getRipsCausaExternaV2(row []string) *models.RipsCausaExternaV2 {
	err := validateVoidData(row, []int{3, 5, 13, 14, 15, 16, 17, 18, 19, 21})
	if err != nil {
		log.Panic(err)
	}
	return &models.RipsCausaExternaV2{
		Code:            row[1],
		Name:            row[2],
		IsAvailable:     row[4] == "SI",
		IsStandartGel:   row[6] == "Verdadero",
		IsStandardMSPS:  row[7] == "Verdadero",
		Consults:        row[8],
		Procedure:       row[9],
		Urgency:         row[10],
		Hospitalization: row[11],
		RnBorn:          row[12],
		UpdateDate:      *updateSisproFormat(row[20]),
	}
}

func getRipsFinalidadConsultaV2(row []string) *models.RipsConsultFinalV2 {
	err := validateVoidData(row, []int{3, 5, 13, 14, 15, 16, 17, 18, 19, 21})
	if err != nil {
		log.Panic(err)
	}
	return &models.RipsConsultFinalV2{
		Code:            row[1],
		Name:            row[2],
		IsAvailable:     row[4] == "SI",
		IsStandartGel:   row[6] == "Verdadero",
		IsStandardMSPS:  row[7] == "Verdadero",
		Consults:        row[8],
		Procedure:       row[9],
		Urgency:         row[10],
		Hospitalization: row[11],
		RnBorn:          row[12],
		UpdateDate:      *updateSisproFormat(row[20]),
	}
}

func getRipsDiagnostictypePrincipalv2(row []string) *models.RipsDiagnostictypePrincipalv2 {
	err := validateVoidData(row, []int{3, 5, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 21})
	if err != nil {
		log.Panic(err)
	}
	return &models.RipsDiagnostictypePrincipalv2{
		Code:           row[1],
		Name:           row[2],
		IsAvailable:    row[4] == "SI",
		IsStandartGel:  row[6] == "Verdadero",
		IsStandardMSPS: row[7] == "Verdadero",
		UpdateDate:     *updateSisproFormat(row[20]),
	}
}

func getService(row []string) *models.Service {
	err := validateVoidData(row, []int{5, 18, 19, 21})
	if err != nil {
		log.Panic(err)
	}
	return &models.Service{
		Code:           row[1],
		Name:           row[2],
		Description:    row[3],
		IsAvailable:    row[4] == "SI",
		IsStandartGel:  row[6] == "Verdadero",
		IsStandardMSPS: row[7] == "Verdadero",
		Extra_I:        row[8],
		Extra_II:       row[9],
		Extra_III:      row[10],
		Extra_IV:       row[11],
		Extra_V:        row[12],
		Extra_VI:       row[13],
		Extra_VII:      row[14],
		Extra_VIII:     row[15],
		Extra_IX:       row[16],
		Extra_X:        row[17],
		UpdateDate:     *updateSisproFormat(row[20]),
	}
}

func getTypeIdPisis(row []string) *models.TypeIdPISIS {
	err := validateVoidData(row, []int{3, 5, 14, 15, 17, 18, 19})
	if err != nil {
		log.Panic(err)
	}
	return &models.TypeIdPISIS{
		Code:            row[1],
		Name:            row[2],
		IsAvailable:     row[4] == "SI",
		IsStandartGel:   row[6] == "Verdadero",
		IsStandardMSPS:  row[7] == "Verdadero",
		ExtraI:          row[8],
		ExtraII:         row[9],
		ExtraIII:        row[10],
		ExtraIV:         row[11],
		ExtraV:          row[12],
		ExtraVI:         row[13],
		ExtraIX:         row[16],
		UpdateDate:      *updateSisproFormat(row[20]),
		IsPublicPrivate: isPublicPrivate(row, 21),
	}
}

func getMedicTypePOS(row []string) *models.MedicTypePOS {
	err := validateVoidData(row, []int{3, 5, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 21})
	if err != nil {
		log.Panic(err)
	}
	return &models.MedicTypePOS{
		Code:           row[1],
		Name:           row[2],
		IsAvailable:    row[4] == "SI",
		IsStandartGel:  row[6] == "Verdadero",
		IsStandardMSPS: row[7] == "Verdadero",
		UpdateDate:     *updateSisproFormat(row[20]),
	}
}

func getTypeNote(row []string) *models.TypeNote {
	err := validateVoidData(row, []int{3, 5, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 21})
	if err != nil {
		log.Panic(err)
	}
	return &models.TypeNote{
		Code:           row[1],
		Name:           row[2],
		IsAvailable:    row[4] == "SI",
		IsStandartGel:  row[6] == "Verdadero",
		IsStandardMSPS: row[7] == "Verdadero",
		UpdateDate:     *updateSisproFormat(row[20]),
	}
}

func getAnotherService(row []string) *models.OtherService {
	err := validateVoidData(row, []int{3, 5, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 21})
	if err != nil {
		log.Panic(err)
	}
	return &models.OtherService{
		Code:           row[1],
		Name:           row[2],
		IsAvailable:    row[4] == "SI",
		IsStandartGel:  row[6] == "Verdadero",
		IsStandardMSPS: row[7] == "Verdadero",
		UpdateDate:     *updateSisproFormat(row[20]),
	}
}
func getUmm(row []string) *models.UMM {
	err := validateVoidData(row, []int{5, 11, 12, 13, 14, 15, 16, 17, 18, 19})
	if err != nil {
		log.Panic(err)
	}
	return &models.UMM{
		Code:            row[1],
		Name:            row[2],
		Description:     row[3],
		IsAvailable:     row[4] == "SI",
		IsStandartGel:   row[6] == "Verdadero",
		IsStandardMSPS:  row[7] == "Verdadero",
		ExtraI:          row[8],
		ExtraII:         row[9],
		ExtraIII:        row[10],
		UpdateDate:      *updateSisproFormat(row[20]),
		IsPublicPrivate: isPublicPrivate(row, 21),
	}
}

func getUpr(row []string) *models.UPR {
	err := validateVoidData(row, []int{3, 5, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19})
	if err != nil {
		log.Panic(err)
	}
	return &models.UPR{
		Code:            row[1],
		Name:            row[2],
		IsAvailable:     row[4] == "SI",
		IsStandartGel:   row[6] == "Verdadero",
		IsStandardMSPS:  row[7] == "Verdadero",
		UpdateDate:      *updateSisproFormat(row[20]),
		IsPublicPrivate: isPublicPrivate(row, 21),
	}
}

func getIngressUser(row []string) *models.IngressUser {
	err := validateVoidData(row, []int{3, 5, 13, 14, 15, 16, 17, 18, 19, 21})
	if err != nil {
		log.Panic(err)
	}
	return &models.IngressUser{
		Code:            row[1],
		Name:            row[2],
		IsAvailable:     row[4] == "SI",
		IsStandartGel:   row[6] == "Verdadero",
		IsStandardMSPS:  row[7] == "Verdadero",
		Consult:         row[8],
		Procedure:       row[9],
		Emergency:       row[10],
		Hospitalization: row[11],
		RBorn:           row[12],
		UpdateDate:      *updateSisproFormat(row[20]),
	}
}

func getCountry(row []string) *models.Country {
	return &models.Country{
		Code:    row[1],
		Name:    row[2],
		ExtraI:  row[9],
		ExtraII: row[10],
	}
}
