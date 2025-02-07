package storage

import (
    "context"
    "log"

    "github.com/minio/minio-go/v7"
    "github.com/minio/minio-go/v7/pkg/credentials"
)

var MinioClient *minio.Client

func InitMinio() {
    endpoint := "localhost:9000"
    accessKeyID := "minioadmin"
    secretAccessKey := "minioadmin"
    useSSL := false

    client, err := minio.New(endpoint, &minio.Options{
        Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
        Secure: useSSL,
    })
    if err != nil {
        log.Fatalf("Failed to initialize Minio: %v", err)
    }

    MinioClient = client

    // Create a bucket
    bucketName := "photos"
    location := "us-east-1"
    err = MinioClient.MakeBucket(context.Background(), bucketName, minio.MakeBucketOptions{Region: location})
    if err != nil {
        exists, errBucketExists := MinioClient.BucketExists(context.Background(), bucketName)
        if errBucketExists == nil && exists {
            log.Printf("We already own %s\n", bucketName)
        } else {
            log.Fatalf("Failed to create bucket: %v", err)
        }
    }
}