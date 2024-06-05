package rest

type DatosGov struct {
	Region           string `json:"__select_alias__"`
	CodeSubregion    string `json:"__select_alias1__"`
	SubRegion        string `json:"__select_alias2__"`
	CodeMunicipality string `json:"__select_alias3__"`
	Municipality     string `json:"__select_alias4__"`
}
