package minio_test

import (
	"context"
	"errors"
	"os"
	"testing"

	minio "github.com/aria-afk/go-get-tickets/minio"
	"github.com/aria-afk/go-get-tickets/utils"
	min "github.com/minio/minio-go/v7"
	qrcode "github.com/skip2/go-qrcode"
)

var m minio.Minio

func TestConnection(t *testing.T) {
	ctx := context.Background()
	utils.LoadEnv("../dev.env")
	client, err := minio.NewMinio()
	if err != nil {
		t.Error(err)
	}
	m = client

	// also ensure buckets are made
	buckets, err := m.Client.ListBuckets(ctx)
	if err != nil {
		t.Error(err)
	}
	bucketList := map[string]bool{}
	for _, bucket := range buckets {
		bucketList[bucket.Name] = true
	}
	_, test := bucketList["qrcodes"]
	_, test1 := bucketList["posters"]
	if !test || !test1 {
		t.Fail()
	}
}

func TestUpload(t *testing.T) {
	ctx := context.Background()
	qrcode.WriteFile("uuid", qrcode.High, 256, "/tmp/uuid.png")
	qrcode.WriteFile("uuid1", qrcode.High, 256, "/tmp/uuid1.png")

	// can upload and not delete local copy
	err := m.UploadImage(ctx, "qrcodes", "uuid", "/tmp/uuid.png", false)
	if err != nil {
		t.Error(err)
	}
	if _, err := os.Stat("/tmp/uuid.png"); errors.Is(err, os.ErrNotExist) {
		t.Error(err)
	}
	// can upload and delete local copy
	err = m.UploadImage(ctx, "qrcodes", "uuid1", "/tmp/uuid1.png", true)
	if err != nil {
		t.Error(err)
	}
	if _, err := os.Stat("/tmp/uuid1.png"); !errors.Is(err, os.ErrNotExist) {
		t.Fail()
	}

	// ensure we can reach the file in the bucket as well.
	file, err := m.Client.GetObject(ctx, "qrcodes", "uuid", min.GetObjectOptions{})
	obj, _ := file.Stat()
	if obj.Key != "uuid" {
		t.Fail()
	}
	file, err = m.Client.GetObject(ctx, "qrcodes", "uuid1", min.GetObjectOptions{})
	obj, _ = file.Stat()
	if obj.Key != "uuid1" {
		t.Fail()
	}
}
