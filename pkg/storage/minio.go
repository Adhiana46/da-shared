package storage

import (
	"context"
	"io"
	"time"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/pkg/errors"
)

type MinioStorage struct {
	client     *minio.Client
	bucketName string
}

// NewMinioStorage initializes a new MinioStorage instance.
func NewMinioStorage(ctx context.Context, endpoint, accessKey, secretKey, bucketName string, useSSL bool) (IStorage, error) {
	client, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKey, secretKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		return nil, errors.Wrap(err, "error creating MinIO client")
	}

	return &MinioStorage{
		client:     client,
		bucketName: bucketName,
	}, nil
}

// Put uploads a file to MinIO.
func (s *MinioStorage) Put(ctx context.Context, file io.Reader, filename string) error {
	_, err := s.client.PutObject(ctx, s.bucketName, filename, file, -1, minio.PutObjectOptions{})
	if err != nil {
		return errors.Wrap(err, "error uploading file to MinIO bucket")
	}

	return nil
}

// Get retrieves a file from MinIO.
func (s *MinioStorage) Get(ctx context.Context, filename string) (io.ReadCloser, error) {
	object, err := s.client.GetObject(ctx, s.bucketName, filename, minio.GetObjectOptions{})
	if err != nil {
		return nil, errors.Wrap(err, "error getting file from MinIO")
	}

	return object, nil
}

// GetLink generates a signed URL for accessing the file.
func (s *MinioStorage) GetLink(ctx context.Context, filename string, duration time.Duration) (FileLink, error) {
	signedURL, err := s.client.PresignedGetObject(ctx, s.bucketName, filename, duration, nil)
	if err != nil {
		return "", errors.Wrap(err, "error generating signed URL for file")
	}

	return FileLink(signedURL.String()), nil
}

// Destroy deletes a file from MinIO.
func (s *MinioStorage) Destroy(ctx context.Context, filename string) error {
	err := s.client.RemoveObject(ctx, s.bucketName, filename, minio.RemoveObjectOptions{})
	if err != nil {
		return errors.Wrap(err, "error deleting file from MinIO")
	}
	return nil
}

// Close closes the MinIO client.
func (s *MinioStorage) Close(ctx context.Context) error {
	// No explicit close method for MinIO client, so this can be a no-op
	return nil
}
