package main

import (
	"github.com/jSierraB3991/scripts/21-ignors/cmd"
	"github.com/jSierraB3991/scripts/21-ignors/configuration"
)

func main() {

	templates := configuration.Configuration("./", "ignors").Templates
	cmd.Run(templates)
}
