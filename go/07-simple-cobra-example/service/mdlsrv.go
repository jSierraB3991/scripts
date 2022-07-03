package service

func (t *Srv) CreateMdl(destFileName string, verFileName string) {
	t.CreateDefaultGoFile("models", verFileName, "models", destFileName)
}

func (t *Srv) CreateMdlDefault(destFileName string) {
	t.CreateMdl(destFileName, "default")
}
func (t *Srv) CreateMdlCrud(destFileName string) {
	t.CreateMdl(destFileName, "crud")
}
