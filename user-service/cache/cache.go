package cache

import "context"

type Cache interface {
	Get(key string, ctx context.Context) (interface{}, error)
	Set(key string, value interface{}, ctx context.Context) error
}
