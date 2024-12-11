// utils/env.go: Tools to load or manage .env files
package utils

import (
	"github.com/joho/godotenv"
)

// Load a specified .env file by specified path (relative)
func LoadEnv(filePath string) error {
	err := godotenv.Load(filePath)
	if err != nil {
		return err
	}
	return nil
}
