package pg

import (
	"os"
	"testing"

	"github.com/aria-afk/go-get-tickets/utils"
)

var pg *PG

func TestNewPG(t *testing.T) {
	os.WriteFile("pg_test.env", []byte("PG_CONN_STRING=postgresql://postgres:test@localhost/template1"), 0755)
	utils.LoadEnv("pg_test.env")

	// Open connection
	db, err := NewPG()
	if err != nil {
		os.Remove("pg_test.env")
		t.Fatalf("Error creating new PG instance:\n%s", err)
	}

	// Ensure connection can preform queries
	var response int64
	err = db.Conn.QueryRow("SELECT 1 + $1", 1).Scan(&response)

	if response != 2 || err != nil {
		os.Remove("pg_test.env")
		t.Fatalf("Error preforming simple query:\n%s", err)
	}

	pg = db

	os.Remove("pg_test.env")
}

func TestLoadQueryMap(t *testing.T) {
	os.Mkdir("queries_test", 0755)
	os.WriteFile("queries_test/one.sql", []byte("SELECT 1 + $1;"), 0755)
	os.Mkdir("queries_test/nested", 0755)
	os.WriteFile("queries_test/nested/two.sql", []byte("SELECT $1 + $2;"), 0755)

	pg.LoadQueryMap("queries_test")

	one, ok := pg.QueryMap["queries_test/one"]
	if ok {
		if one != "SELECT 1 + $1;" {
			t.Fatalf("Query map loaded query_test/one.sql but did not parse its contents properly")
		}
	} else {
		t.Fatalf("Query map did not load singly nested query queries_test/one.sql")
	}

	two, ok := pg.QueryMap["queries_test/nested/two"]
	if ok {
		if two != "SELECT $1 + $2;" {
			t.Fatalf("Query map loaded query_test/nested/two.sql but did not parse its contents properly")
		}
	} else {
		t.Fatalf("Query map did not load singly nested query queries_test/nested/two.sql")
	}

	os.RemoveAll("queries_test")
}
