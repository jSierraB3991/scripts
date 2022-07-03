package main

import (
	"github.com/OkabeRitarou/11-gorm-toml/libs"
)

func main() {
	dbConfig := libs.Configure("./", "mysql")
	dbConfig.InitMysqlDB()
}
