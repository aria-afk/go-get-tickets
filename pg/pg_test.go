package pg

import (
	"os"
	"testing"

	"github.com/aria-afk/go-get-tickets/queries"
	"github.com/aria-afk/go-get-tickets/utils"
)

func setup() {
	os.WriteFile("pg_test.env", []byte("PG_CONN_STRING=postgresql://postgres:test@localhost/template1?sslmode=disable"), 0755)
	utils.LoadEnv("pg_test.env")
}

func cleanup() {
	os.Remove("pg_test.env")
}

func TestNewPG(t *testing.T) {
	setup()

	// Open connection
	db, err := NewPG()
	if err != nil {
		cleanup()
		t.Fatalf("Error creating new PG instance:\n%s", err)
	}

	// Ensure connection can perform queries
	var response int
	err = db.Conn.QueryRow(queries.TestQuery, 1).Scan(&response)

	if response != 1 || err != nil {
		cleanup()
		t.Fatalf("Error preforming simple query:\n%s", err)
	}

	cleanup()
}

func TestMigrations(t *testing.T) {
	db, err := NewPG()
	if err != nil {
		cleanup()
		t.Fatalf("Error creating new PG instance:\n%s", err)
	}
	err = db.Migrate("up")
	if err != nil {
		cleanup()
		t.Fatalf("Error migrating up:\n%s", err)
	}
	err = db.Migrate("down")
	if err != nil {
		cleanup()
		t.Fatalf("Error migrating down:\n%s", err)
	}

	cleanup()
}
