package utils

import (
	"github.com/joho/godotenv"
)

// Load a specified .env file by local path
func LoadEnv(localPath string) error {
	err := godotenv.Load(localPath)
	if err != nil {
		return err
	}
	return nil
}
