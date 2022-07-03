package service

func (t *Srv) CreatePodman(filesVersionName string) {
	origFolderName := "podman"
	desFileName := "Containerfile"
	t.CopyTemplate(origFolderName, filesVersionName, desFileName)
}

func (t *Srv) PodmanGo() {
	t.CreatePodman("podman")
}
