// qr.go tool to generate qr codes for provisioned tickets
package utils

import (
	"context"
	"fmt"

	"github.com/aria-afk/go-get-tickets/minio"

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

func storeQRCode(ticketUUID string) error {
	m, err := minio.NewMinio()
	if err != nil {
		return err
	}
	err = m.UploadImage(context.Background(), "qrcodes", ticketUUID, fmt.Sprintf("/tmp/%s.png", ticketUUID), true)
	if err != nil {
		return err
	}

	return nil
}
