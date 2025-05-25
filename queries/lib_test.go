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

/* Uncomment to test
func TestParseQueryFilePanic(t *testing.T) {
	testPanic := parseQueryFile("./_test/DOESNTEXIST.sql")
	if len(testPanic) > 0 {
		t.Error("Did not panic out on a non-existent .sql file")
	}
}
*/
