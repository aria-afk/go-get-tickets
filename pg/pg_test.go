package pg

import (
	"os"
	"testing"

	"github.com/aria-afk/go-get-tickets/utils"
)

func setup() {
	os.WriteFile("pg_test.env", []byte("PG_CONN_STRING=postgresql://postgres:test@localhost/template1"), 0755)
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

	// Ensure connection can preform queries
	var response int64
	err = db.Conn.QueryRow("SELECT 1 + $1", 1).Scan(&response)

	if response != 2 || err != nil {
		cleanup()
		t.Fatalf("Error preforming simple query:\n%s", err)
	}

	cleanup()
}
