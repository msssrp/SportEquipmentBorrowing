package function

import (
	"github.com/joho/godotenv"
	"os"
)

func GetDotEnv(key string) (string, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return "", err
	}
	return os.Getenv(key), nil
}
