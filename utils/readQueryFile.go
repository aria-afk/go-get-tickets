package utils

import (
	"fmt"
	"os"
)

func ReadQueryFile(path string) string {
	// Ensure file has a .sql extension
	if path[len(path)-4:] != ".sql" {
		panic(fmt.Sprintf("Provided file path is not a .sql file, given path: %s", path))
	}
	data, err := os.ReadFile(path)
	if err != nil {
		panic(fmt.Sprintf("Issue parsing query file [ %s ]\n Error: %s", path, err))
	}
	return string(data)
}
