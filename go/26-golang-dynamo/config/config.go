package config

import (
	"strconv"

	"github.com/jsierrab3991/scripts/26-golang-dynamo/utils/env"
)

type Config struct {
	Port        int
	Timeout     int
	Dialect     string
	DatabaseUrl string
}

func GetConfig() *Config {
	return &Config{
		Port:        parseEnvToInt("PORT", 3000),
		Timeout:     parseEnvToInt("TIMEOUT", 30),
		Dialect:     env.GetEnv("DIALECT", "sqlite3"),
		DatabaseUrl: env.GetEnv("DATABASE_URL", ":memory:"),
	}
}

func parseEnvToInt(enviroment string, defaultValue int) int {
	data, err := strconv.Atoi(env.GetEnv(enviroment, strconv.Itoa(defaultValue)))
	if err != nil {
		return defaultValue
	}
	return data
}
