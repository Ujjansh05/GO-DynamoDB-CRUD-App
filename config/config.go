package config

import(
	"github.com/Ujjansh05/GO_DynamoDB_CRUD_App/utlis/env"
	"strconv"
)

type Config struct{
	Port 		int 
	Timeout 	int
	Dialect		string
	DatabaseURI	string
}

func GetConfig() Config{
	return Config{
		Port:		parseEnvToInt("PORT", "8080") 
		Timeout:	parseEnvToInt("TIMEOUT", "30")
		Dialect: 	env.GetEnv("DIALECT", "sqlite3")
		DatabaseURI: env.GetEnv("DatabaseURI", "memory:")
	}
}

func parseEnvToInt(envName, defaultValue string) int{
	num, err := strconv.Atoi(env.GetEnv(env, defaultValue))
	if err != nil {
		return 0
	}
	return num
}