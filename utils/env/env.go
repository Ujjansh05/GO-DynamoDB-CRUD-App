package env


import "os"

func GetEnv(env, defaultValue string) string{
	enviroment := os.GetEnv(env)
	if enviroment = "" {
		return defaultValue
	}
	return enviroment
}