package queries

import (
	"strings"
	"testing"
)

func TestParseQueryFilePass(t *testing.T) {
	givenQuery := strings.TrimSuffix(testQuery, "\n")
	expectedQuery := "SELECT * FROM users WHERE id = $1;"
	if givenQuery != expectedQuery {
		t.Errorf("Obtaining the query file didnt crash, however the file was not parsed properly.\n Expected:%s\n Recieved:%s", givenQuery, expectedQuery)
	}
}

func TestParseQueryFilePanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("Did not panic out on a non-existent .sql file")
		}
	}()
	_ = parseQueryFile("./_test/DOESNTEXIST.sql")
}
