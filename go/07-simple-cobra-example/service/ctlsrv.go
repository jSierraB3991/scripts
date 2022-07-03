package service

func (t *Srv) CreateCtlDefault(destFileName string) {
	t.CreateDefaultGoFile("controllers", "default", "controllers", destFileName)
}
func (t *Srv) CreateCtlCrud(destFileName string) {
	t.CreateDefaultGoFile("controllers", "crud", "controllers", destFileName)
}
