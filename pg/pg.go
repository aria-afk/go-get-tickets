// pg: sub-package to provide connection to Postgresql as well as some utility methods
package pg

import (
	"database/sql"
	"errors"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

type PG struct {
	conn *sql.DB
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
		conn: db,
	}, nil
}
