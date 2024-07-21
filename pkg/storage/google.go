package storage

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"time"

	"cloud.google.com/go/storage"
	"github.com/pkg/errors"
	"google.golang.org/api/option"
)

const (
	GCS_CRED_TYPE       = "service_account"
	GCS_AUTH_URI        = "https://accounts.google.com/o/oauth2/auth"
	GCS_TOKEN_URI       = "https://oauth2.googleapis.com/token"
	GCS_AUTH_CERT_URL   = "https://www.googleapis.com/oauth2/v1/certs"
	GCS_CLIENT_CERT_URL = "https://www.googleapis.com/robot/v1/metadata/x509/graphical-bus-99503%40appspot.gserviceaccount.com"
)

type GoogleStorageCredential struct {
	ProjectId    string `json:"project_id"`
	PrivateKeyId string `json:"private_key_id"`
	PrivateKey   string `json:"private_key"`
	ClientEmail  string `json:"client_email"`
	ClientId     string `json:"client_id"`
}

type GoogleStorage struct {
	client     *storage.Client
	bucketName string
}

// NewGoogleStorage initializes a new GoogleStorage instance.
func NewGoogleStorage(ctx context.Context, bucketName string, credential GoogleStorageCredential) (IStorage, error) {
	credMap := map[string]any{
		"type":                        GCS_CRED_TYPE,
		"project_id":                  credential.ProjectId,
		"private_key_id":              credential.PrivateKeyId,
		"private_key":                 credential.PrivateKey,
		"client_email":                credential.ClientEmail,
		"client_id":                   credential.ClientId,
		"auth_uri":                    GCS_AUTH_URI,
		"token_uri":                   GCS_TOKEN_URI,
		"auth_provider_x509_cert_url": GCS_AUTH_CERT_URL,
		"client_x509_cert_url":        GCS_CLIENT_CERT_URL,
	}

	credJson, err := json.Marshal(credMap)
	if err != nil {
		return nil, errors.Wrap(err, "error marshaling JSON credential")
	}

	client, err := storage.NewClient(ctx, option.WithCredentialsJSON(credJson))
	if err != nil {
		return nil, errors.Wrap(err, "error creating Google Cloud Storage client")
	}

	return &GoogleStorage{
		client:     client,
		bucketName: bucketName,
	}, nil
}

// Put uploads a file to Google Cloud Storage.
func (s *GoogleStorage) Put(ctx context.Context, file io.Reader, filename string) error {
	writer := s.client.Bucket(s.bucketName).Object(filename).NewWriter(ctx)
	defer func() {
		if err := writer.Close(); err != nil {
			// Log error if closing writer fails
			slog.Error(fmt.Sprintf("error closing writer: %v", err))
		}
	}()

	if _, err := io.Copy(writer, file); err != nil {
		return errors.Wrap(err, "error copying file to Google Cloud Storage bucket")
	}

	return nil
}

// Get retrieves a file from Google Cloud Storage.
func (s *GoogleStorage) Get(ctx context.Context, filename string) (io.ReadCloser, error) {
	reader, err := s.client.Bucket(s.bucketName).Object(filename).NewReader(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "error getting file from Google Cloud Storage")
	}

	return reader, nil
}

// GetLink generates a signed URL for accessing the file.
func (s *GoogleStorage) GetLink(ctx context.Context, filename string, duration time.Duration) (FileLink, error) {
	opts := &storage.SignedURLOptions{
		Scheme:  storage.SigningSchemeV4,
		Method:  "GET",
		Expires: time.Now().Add(duration),
	}

	signedURL, err := s.client.Bucket(s.bucketName).SignedURL(filename, opts)
	if err != nil {
		return "", errors.Wrap(err, "error generating signed URL for file")
	}

	return FileLink(signedURL), nil
}

// Destroy deletes a file from Google Cloud Storage.
func (s *GoogleStorage) Destroy(ctx context.Context, filename string) error {
	obj := s.client.Bucket(s.bucketName).Object(filename)
	if err := obj.Delete(ctx); err != nil {
		return errors.Wrap(err, "error deleting file from Google Cloud Storage")
	}
	return nil
}

// Close closes the Google Cloud Storage client.
func (s *GoogleStorage) Close(ctx context.Context) error {
	if err := s.client.Close(); err != nil {
		return errors.Wrap(err, "error closing Google Cloud Storage client")
	}
	return nil
}
