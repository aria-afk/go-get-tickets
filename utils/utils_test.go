package utils

import (
	"os"
	"testing"
)

func TestLoadEnv(t *testing.T) {
	os.WriteFile("utils_test.env", []byte("TEST=1337"), 0755)

	err := LoadEnv("./utils_test.env")
	if err != nil {
		os.Remove("utils_test.env")
		t.Fatalf("Error loading utils_test.env - %s", err.Error())
	}

	test := os.Getenv("TEST")
	if test != "1337" {
		os.Remove("utils_test.env")
		t.Fatalf("Error accessing TEST from utils_test.env - %s", err.Error())
	}

	os.Remove("utils_test.env")
}
