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

type IPSCodHabilitacionRepositoryInterface interface {
	SaveIPSCodHabilitacion(data *models.IpsCpdeHabilitation) error
}

type IpsNoRepsRepositoryInterface interface {
	SaveIpsNoReps(data *models.IpsNoReps) error
}

type IumsRepositoryInterface interface {
	SaveIum(data *models.Ium) error
}

type ModalityAtentionRepositoryInterface interface {
	SaveAtentionModality(data *models.AtentionModality) error
}

type RipsExternCauseRepositoryInterface interface {
	SaveRipsExternCauseV2(data *models.RipsCausaExternaV2) error
}

type RipsConsultFinalRepositoryInterface interface {
	SaveRipsConsultFinalV2(data *models.RipsConsultFinalV2) error
}

type RipsDiagnosticTypePincipalRepositoryInterface interface {
	SaveRipsDiagnostictypePrincipalv2(data *models.RipsDiagnostictypePrincipalv2) error
}

type UserTypeRepositoryInterface interface {
	SaveUserType(data *models.UserType) error
}

type ServiceRepositoryInterface interface {
	SaveService(data *models.Service) error
}

type OtherServiceRepositoryInterface interface {
	SaveOtherService(data *models.OtherService) error
}

type TypeIdPISISRepositoryInterface interface {
	SaveTypeIdPISIS(data *models.TypeIdPISIS) error
}

type MedicTypePOSRepositoryInterface interface {
	SaveMedicTypePOS(data *models.MedicTypePOS) error
}

type TypeNotePOSRepositoryInterface interface {
	SaveTypeNote(data *models.TypeNote) error
}

type UMMRepositoryInterface interface {
	SaveUMM(data *models.UMM) error
}

type UPRRepositoryInterface interface {
	SaveUPR(data *models.UPR) error
}

type IngressUserRepositoryInterface interface {
	SaveIngressUser(data *models.IngressUser) error
}
