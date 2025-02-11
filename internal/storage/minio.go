package storage

import (
    "context"
    "log"
    "fmt"

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

func UploadFile(bucketName, objectName, filePath, contentType string) error {
    // Upload the file
    _, err := MinioClient.FPutObject(context.Background(), bucketName, objectName, filePath, minio.PutObjectOptions{ContentType: contentType})
    if err != nil {
        return fmt.Errorf("failed to upload file: %v", err)
    }
    return nil
}

func DeleteFile(bucketName, objectName string) error {
    err := MinioClient.RemoveObject(context.Background(), bucketName, objectName, minio.RemoveObjectOptions{})
    if err != nil {
        return fmt.Errorf("failed to delete file: %v", err)
    }
    return nil
}