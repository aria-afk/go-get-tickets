// pg: sub-package to provide connection to Postgresql as well as some utility methods
package pg

import (
	"database/sql"
	"errors"
	"fmt"
	"os"
	"strings"

	_ "github.com/lib/pq"
)

type PG struct {
	Conn     *sql.DB
	QueryMap map[string]string
}

// Constructor to recieve a new *PG instance. Attempts to open a connection to
// postgres via the PG_CONN_STRING in the .env file, assuming it is already loaded.
func NewPG() (*PG, error) {
	connStr := os.Getenv("PG_CONN_STRING")
	if connStr == "" {
		return &PG{}, errors.New("PG ERROR: Could not retrieve the PG_CONN_STRING from .env")
	}

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return &PG{}, fmt.Errorf("PG ERROR: Could not open connection to postgres \n %s", err)
	}

	return &PG{
		Conn:     db,
		QueryMap: make(map[string]string, 0),
	}, nil
}

// Parses a given base dirPath and stores all query files into
// PG.QueryMap with the format: map[absoluteFilePath]queryFileContents.
// example: queries/getVendor = "SELECT * FROM vendors WHERE name = $1;"
//
// Works recursively through the tree. a nested file may look like:
// queries/tickets/delete = "DELETE FROM tickets WHERE uuid = $1;"
func (pg *PG) LoadQueryMap(dirPath string) error {
	dirs := []string{dirPath}
	for len(dirs) > 0 {
		dir := pop(&dirs)
		files, err := os.ReadDir(dir)
		if err != nil {
			return err
		}

		for _, file := range files {
			path := fmt.Sprintf("%s/%s", dir, file.Name())

			if file.IsDir() {
				dirs = append(dirs, path)
				continue
			}

			data, err := os.ReadFile(path)
			if err != nil {
				return err
			}

			extension := strings.Split(file.Name(), ".")
			// Ensure file type (we may want to remove this)
			if extension[len(extension)-1] != "sql" {
				continue
			}

			pg.QueryMap[dir+"/"+extension[0]] = string(data)
		}
	}
	return nil
}

func pop(arr *[]string) string {
	l := len(*arr)
	rv := (*arr)[l-1]
	*arr = (*arr)[:l-1]
	return rv
}
