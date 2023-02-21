package template

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/jSierraB3991/scripts/21-ignors/configuration"
	"github.com/schollz/progressbar/v3"
)

func DownloadIgnor(language string) {
	home := os.Getenv("HOME")
	config := configuration.Configuration(home+"/.config/ignors", "ignors")
	url := config.Url

	if !strings.Contains(strings.Join(config.Templates, ","), language) {
		fmt.Println("The language: ", language, " is not present on languages list")
		os.Exit(1)
	}

	definitiveUrl := fmt.Sprintf("%s/%s.gitignore", url, language)
	req, _ := http.NewRequest("GET", definitiveUrl, nil)
	resp, _ := http.DefaultClient.Do(req)
	defer resp.Body.Close()

	bar := progressbar.DefaultBytes(
		resp.ContentLength,
		"downloading",
	)
	filepath := strings.Join([]string{".gitignore"}, "/")
	out, err := os.Create(filepath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer out.Close()
	_, err = io.Copy(io.MultiWriter(out, bar), resp.Body)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
