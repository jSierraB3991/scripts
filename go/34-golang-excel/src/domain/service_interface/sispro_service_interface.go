package serviceinterface

type SisproServiceInterface interface {
	SaveSisproData(data interface{}) error
	GetCodesForData() ([]string, error)
}
