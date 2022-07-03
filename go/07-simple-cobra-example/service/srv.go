package service

import (
	"fmt"
	"strings"

	"github.com/culturadevops/jgt/jio"
)

var VarSrv *Srv

type Srv struct {
}

func (t *Srv) aUpper(b string) string {
	b = strings.ToLower(b)
	b = strings.Title(b)
	return b
}

func (t *Srv) CreateMapDefault(fileName string) map[string]string {
	MapForReplace := make(map[string]string)
	MapForReplace["%MODELNAME%"] = strings.ToLower(fileName)
	MapForReplace["%EXPORTNAME%"] = t.aUpper(fileName)
	return MapForReplace
}

func (t *Srv) GetTemplateName(origFolderName string, filesVersionName string) string {
	folderProject := "/home/lelouch/Source/scripts/go/07-simple-cobra-example/templates/echo"
	fmt.Println(folderProject + "/" + origFolderName + "/" + filesVersionName + ".stub")
	return folderProject + "/" + origFolderName + "/" + filesVersionName + ".stub"
}

func (t *Srv) CreateDefaultFile(origFolderName string, filesVersionName string, destFolderName string, nameStruct string, destFileName string) {
	jio.CreateFolder(origFolderName)
	MapForReplace := t.CreateMapDefault(nameStruct)
	newName := destFolderName + "/" + destFileName
	jio.NewFileforTemplate(newName, t.GetTemplateName(origFolderName, filesVersionName), MapForReplace)
}

func (t *Srv) CreateDefaultGoFile(origFolderName string, filesVersionName string, destFolderName string, destFileName string) {
	nameStruct := destFileName
	destFileName = destFileName + ".go"
	t.CreateDefaultFile(origFolderName, filesVersionName, destFolderName, nameStruct, destFileName)
}

func (t *Srv) CopyTemplate(origFolderName string, filesVersionName string, destFileName string) {
	jio.Copy(t.GetTemplateName(origFolderName, filesVersionName), destFileName)
}
