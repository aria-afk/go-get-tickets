// qr.go tool to generate qr codes for provisioned tickets
package utils

import (
	"fmt"

	go_qr "github.com/piglig/go-qr"
)

func GenerateQRCode(ticketUUID string) error {
	errCorLvl := go_qr.High
	qr, err := go_qr.EncodeText(ticketUUID, errCorLvl)
	if err != nil {
		return err
	}
	config := go_qr.NewQrCodeImgConfig(10, 4)
	err = qr.PNG(config, fmt.Sprintf("/tmp/%s.png", ticketUUID))
	if err != nil {
		return err
	}

	return nil
}

func storeQRCode(path string) error {
	return nil
}

// TODO:
func GetQRCode(path string) error {
	return nil
}

// TODO:
// does min.io do expire times?
// otherwise we can run a job or something to clean it up
func DeleteQRCode(path string) error {
	return nil
}
