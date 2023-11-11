package utilities

import (
	"os"

	"github.com/joho/godotenv"
)

func LoadOnlyGetEnv(key string) string {
	return os.Getenv(key)
}

func GoDotEnvVariableTest(key string) string {

	// load .env file
	err := godotenv.Load("../.env")

	if err != nil {
		return LoadOnlyGetEnv(key)
	}

	return os.Getenv(key)
}

func GoDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		return GoDotEnvVariableTest(key)
	}

	return os.Getenv(key)
}
