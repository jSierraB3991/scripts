package service

func (t *Srv) CreateMid(destFileName string) {
	t.CreateDefaultGoFile("middlewares", "default", "middlewares", destFileName)
}
