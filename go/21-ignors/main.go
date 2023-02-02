package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/jSierraB3991/scripts/21-ignors/cmd"
	"github.com/jSierraB3991/scripts/21-ignors/configuration"
	"github.com/jSierraB3991/scripts/21-ignors/template"
)

func questionLanguages() {
	home := os.Getenv("HOME")
	templates := configuration.Configuration(home+"/.config/ignors", "ignors").Templates
	cmd.Run(templates)
}

func main() {
	languagesCmd := flag.String("lan", "", "languages for download ignor")
	if len(os.Args) < 2 {
		questionLanguages()
	} else {
		flag.Parse()
		if *languagesCmd == "" {
			fmt.Println("if you konw language for ignor -lan LANGUAGE, example : -lan Java")
			os.Exit(1)
		}
		template.DownloadIgnor(*languagesCmd)
	}
}
