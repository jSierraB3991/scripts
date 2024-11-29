package repositoryinterface

import "github.com/jdsierrab3991/scripts/34-golang-excel/domain/models"

type CieRepositoryInterface interface {
	SaveCie2036SisPro(data *models.Cie2036) error
	SaveCieSisPro(data *models.Cie) error
	GetCodesForCie() ([]string, error)
	GetCodesForCie2036() ([]string, error)
}

type CumRepositoryInterface interface {
	SaveCumSisPro(data *models.CumSispro) error
	GetCodesForDataSispro() ([]string, error)
}

type CollectionConceptRepositoryInterface interface {
	SaveCollectionConceptSisPro(data *models.CollectionConcept) error
	GetCodesForCollectionConcept() ([]string, error)
}

type CupRipsRepositoryInterface interface {
	SaveCupRips(data *models.CupsRips) error
	GetCodesForCupsRips() ([]string, error)
}

type DcirepositoryInterface interface {
	SaveDci(data *models.Dci) error
	GetCodesForDci() ([]string, error)
}

type UserEgreRepositoryInterface interface {
	SaveUserEgrese(data *models.UserEgrese) error
	GetCodesForUserEgrese() ([]string, error)
}

type FfmRepositoryInterface interface {
	SaveFfm(data *models.Ffm) error
	GetCodesForFfm() ([]string, error)
}

type GroupServiceRepositoryInterface interface {
	SaveGroupService(data *models.GroupService) error
	GetCodesForDataGroupService() ([]string, error)
}

type IPSCodHabilitacionRepositoryInterface interface {
	SaveIPSCodHabilitacion(data *models.IpsCpdeHabilitation) error
	GetCodesForIpsCodeNoHabilitation() ([]string, error)
}

type IpsNoRepsRepositoryInterface interface {
	SaveIpsNoReps(data *models.IpsNoReps) error
	GetCodesForIpsNoReps() ([]string, error)
}

type IumsRepositoryInterface interface {
	SaveIum(data *models.Ium) error
	GetCodesForIum() ([]string, error)
}

type ModalityAtentionRepositoryInterface interface {
	SaveAtentionModality(data *models.AtentionModality) error
	GetCodesForAtentionModality() ([]string, error)
}

type RipsExternCauseRepositoryInterface interface {
	SaveRipsExternCauseV2(data *models.RipsCausaExternaV2) error
	GetCodesForRipsExternCauseV2() ([]string, error)
}

type RipsConsultFinalRepositoryInterface interface {
	SaveRipsConsultFinalV2(data *models.RipsConsultFinalV2) error
	GetCodesForConsultFinalV2() ([]string, error)
}

type RipsDiagnosticTypePincipalRepositoryInterface interface {
	SaveRipsDiagnostictypePrincipalv2(data *models.RipsDiagnostictypePrincipalv2) error
	GetCodesForRipsDiagnosticPPalV2() ([]string, error)
}

type UserTypeRepositoryInterface interface {
	SaveUserType(data *models.UserType) error
	GetCodesForDataUserType() ([]string, error)
}

type ServiceRepositoryInterface interface {
	SaveService(data *models.Service) error
	GetCodesForService() ([]string, error)
}

type OtherServiceRepositoryInterface interface {
	SaveOtherService(data *models.OtherService) error
	GetCodesForDataOtherService() ([]string, error)
}

type TypeIdPISISRepositoryInterface interface {
	SaveTypeIdPISIS(data *models.TypeIdPISIS) error
	GetCodesForTypeIdPisis() ([]string, error)
}

type MedicTypePOSRepositoryInterface interface {
	SaveMedicTypePOS(data *models.MedicTypePOS) error
	GetCodesForMedicTypeOs() ([]string, error)
}

type TypeNotePOSRepositoryInterface interface {
	SaveTypeNote(data *models.TypeNote) error
	GetCodesForTypeNote() ([]string, error)
}

type UMMRepositoryInterface interface {
	SaveUMM(data *models.UMM) error
	GetCodesForDataUmm() ([]string, error)
}

type UPRRepositoryInterface interface {
	SaveUPR(data *models.UPR) error
	GetCodesForDataUpr() ([]string, error)
}

type IngressUserRepositoryInterface interface {
	SaveIngressUser(data *models.IngressUser) error
	GetCodesForDataIngresUser() ([]string, error)
}
type CountryRepositoryInterface interface {
	SaveCountry(data *models.Country) error
	GetCodesForCountry() ([]string, error)
}
