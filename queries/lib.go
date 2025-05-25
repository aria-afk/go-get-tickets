package queries

import (
	"fmt"
	"os"
	"path/filepath"
)

func parseQueryFile(path string) string {
	// Ensure file is a valid .sql file
	if path[len(path)-4:] != ".sql" {
		panic(fmt.Sprintf("Provided file path is not a .sql file, given path: %s", path))
	}
	data, err := os.ReadFile(path)
	if err != nil {
		panic(fmt.Sprintf("Issue parsing query file [ %s ]\n Error: %s", path, err))
	}
	return string(data)
}

// Ensure this package can be imported properly from any supported sub-package
func formatBasePath() string {
	absPath, err := filepath.Abs("./")
	if err != nil {
		panic(err)
	}
	basePath := filepath.Base(absPath)
	switch basePath {
	// Root dir
	case "go-get-tickets":
		return filepath.Join(absPath, "queries")
	// This package (for queries_test)
	case "queries":
		return absPath
	default:
		panic("Given working dir does not have a mapping yet, please add it.")
	}
}

var basePath = formatBasePath()
