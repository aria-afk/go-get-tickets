// pg: sub-package to provide connection to Postgresql as well as some utility methods
package pg

import (
	"database/sql"
	"errors"
	"fmt"
	"os"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

type PG struct {
	Conn *sql.DB
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

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		fmt.Println("migrate error 1")
		fmt.Print(err)
	}
	fmt.Println("hi1")

	m, err := migrate.NewWithDatabaseInstance("file://migrations", "postgres", driver)
	if err != nil {
		fmt.Println("migrate error 2")
		fmt.Print(err)
	}

	// TODO should we fail completely here like this on a failure to apply migrations?
	err = m.Up()
	if err != nil {
		fmt.Printf("Failed to apply migrations:\n %s", err)
		return nil, err
	}

	return &PG{
		Conn: db,
	}, nil
}
