package env

import "os"

func GetEnv(env, defaultValue string) string {
	enviroment := os.Getenv(env)
	if enviroment == "" {
		return defaultValue
	}
	return enviroment
}
