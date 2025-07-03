package minio

import (
	"fmt"
	"os"
	"strings"

	min "github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"golang.org/x/net/context"
)

type Minio struct {
	Client        *min.Client
	DefaultBucket string
}

func NewMinio() (Minio, error) {
	ctx := context.Background()
	endpoint := os.Getenv("MINIO_ENDPOINT")
	accessKeyId := os.Getenv("MINIO_ACCESS_KEY_ID")
	secretAccessKey := os.Getenv("MINIO_SECRET_ACCESS_KEY")
	useSSL := false

	minioClient, err := min.New(endpoint, &min.Options{
		Creds:  credentials.NewStaticV4(accessKeyId, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		return Minio{}, err
	}

	m := Minio{Client: minioClient}
	bucketList := os.Getenv("MINIO_BUCKET_LIST")
	for _, bucketName := range strings.Split(bucketList, ",") {
		err = m.Client.MakeBucket(ctx, bucketName, min.MakeBucketOptions{})
	}
	return m, nil
}

func (m *Minio) upsertBucket(ctx context.Context, bucketName string) error {
	fmt.Println(bucketName)
	err := m.Client.MakeBucket(ctx, bucketName, min.MakeBucketOptions{})
	if err != nil {
		_, alreadyExists := m.Client.BucketExists(ctx, bucketName)
		if alreadyExists == nil {
			return err
		}
	}
	return nil
}

func (m *Minio) UploadImage(ctx context.Context, bucket string, name string, filepath string, delete bool) error {
	_, err := m.Client.FPutObject(ctx, bucket, name, filepath, min.PutObjectOptions{})
	if err != nil {
		return err
	}

	if delete {
		return os.Remove(filepath)
	}
	return nil
}
