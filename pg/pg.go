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

	return &PG{
		Conn: db,
	}, nil
}

func (pg *PG) Migrate(migrationType string) error {
	driver, err := postgres.WithInstance(pg.Conn, &postgres.Config{})
	if err != nil {
		fmt.Printf("Migration error getting driver:\n%s\n", err)
		return err
	}

	m, err := migrate.NewWithDatabaseInstance("file://pg/migrations", "postgres", driver)
	if err != nil {
		fmt.Printf("Migration error getting migrator:\n%s\n", err)
		return err
	}

	switch migrationType {
	case "up":
		err = m.Up()
	case "down":
		err = m.Down()
	default:
		err = errors.New(fmt.Sprintf("Migration error: Invalid migration type provided: %s", migrationType))
	}

	if err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			fmt.Println("No change")
		} else {
			fmt.Printf("Failed to apply migrations:\n%s\n", err)
			return err
		}
	}

	fmt.Println("Successfully applied migrations")
	return nil
}
