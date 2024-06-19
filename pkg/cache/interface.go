package cache

import "errors"

var (
	ErrCacheNil = errors.New("nil")
)

type Cache interface {
	Set(key string, value string, expSecond int32) error
	Get(key string) (string, error)
	GetTTL(key string) (int32, error)
	Delete(key string) error
	Close() error
}
