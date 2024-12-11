// sub-package to provide a base connection
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

// NewPG: Creates a new PG struct instance and attempts to connect
// to the provided psql connection string via ENV (PG_CONN_STRING)
func NewPG() (*PG, error) {
	connStr := os.Getenv("PG_CONN_STRING")
	if len(connStr) == 0 {
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
