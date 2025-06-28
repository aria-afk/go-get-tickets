package utils

import (
	"os"
	"testing"
)

func TestGenerateQRCode(t *testing.T) {
	err := GenerateQRCode("uuid")
	if err != nil {
		t.Error(err)
	}

	if _, err := os.Stat("/tmp/uuid.png"); err != nil {
		t.Error("QR code did not get generated")
	} else {
		os.Remove("/tmp/uuid.png")
	}
}
