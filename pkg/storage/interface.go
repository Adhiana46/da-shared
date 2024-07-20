package storage

import (
	"context"
	"io"
	"time"
)

type FileLink string

type IStorage interface {
	Put(ctx context.Context, file io.Reader, filename string) error
	Get(ctx context.Context, filename string) (io.ReadCloser, error)
	GetLink(ctx context.Context, filename string, duration time.Duration) (FileLink, error)
	Destroy(ctx context.Context, filename string) error
	Close(ctx context.Context) error
}
