package function

import (
	"github.com/joho/godotenv"
	"os"
	"path/filepath"
)

func GetDotEnv(key string) (string, error) {
	// Get the full path to the root folder
	rootDir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	// Load environment variables from .env file
	err = godotenv.Load(filepath.Join(rootDir, ".env"))
	if err != nil {
		return "", err
	}

	return os.Getenv(key), nil
}
