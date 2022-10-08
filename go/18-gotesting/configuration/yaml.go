package configuration

import (
	"io/ioutil"

	"gitlab.com/eliotandelon/gotesting/models"
	"gopkg.in/yaml.v2"
)

func LoadYml(fileName string) (*models.Configuration, error) {
	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	var config = &models.Configuration{}

	err = yaml.Unmarshal(content, config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
