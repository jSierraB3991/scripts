package configuration

import (
	"fmt"
	"log"
	"os"

	"github.com/jSierraB3991/scripts/21-ignors/model"
	config "github.com/spf13/viper"
)

func Configuration(path string, tomlName string) *model.Config {
	config.AddConfigPath(path)
	config.SetConfigName(tomlName)
	if err := config.ReadInConfig(); err != nil {
		log.Fatalf("Error reading viper file, %s", err)
	}
	if config.GetString("url") == "" {
		fmt.Println("no found url in the file " + tomlName)
		os.Exit(1)
	}

	var templates []string
	err := config.UnmarshalKey("templates", &templates)
	if err != nil {
		fmt.Printf("not found templates in the file %s, %v", tomlName, err)
		os.Exit(1)
	}

	response := &model.Config{
		Url:       config.GetString("url"),
		Templates: templates,
	}
	return response
}
