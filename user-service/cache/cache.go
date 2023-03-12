package cache

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type Cache interface {
	Get(key string, ctx context.Context) (interface{}, error)
	Set(key string, value interface{}, ctx context.Context) error
}

type UserCache struct {
	client  *redis.Client
	host    string
	port    string
	db      int
	expires time.Duration
}

func NewRedisCache(host string, port string, db int, expires time.Duration) *UserCache {
	return &UserCache{
		client:  redis.NewClient(&redis.Options{Addr: host + ":" + port}),
		host:    host,
		port:    port,
		db:      db,
		expires: expires,
	}
}
