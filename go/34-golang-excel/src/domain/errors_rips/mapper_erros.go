package errorsrips

type MapperError struct {
	Data string
}

func (mp MapperError) Error() string {
	return "MAPPER_ERROR_" + mp.Data
}

type ValidataInMapperSisPro struct {
	Data string
}

func (mp ValidataInMapperSisPro) Error() string {
	return "SISPRO_VALID_DATA_" + mp.Data
}
