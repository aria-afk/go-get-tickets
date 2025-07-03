package utils

import (
	"testing"
)

func TestQRCodeCreation(t *testing.T) {
	err := GenerateQRCode("qr1")
	if err != nil {
		t.Error(err)
	}
}

func TestQRCodeStorage(t *testing.T) {
	LoadEnv("../dev.env")
	err := storeQRCode("qr1")
	if err != nil {
		t.Error(err)
	}
}
