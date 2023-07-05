package cache

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type Cache interface {
	GetRole(key string, ctx context.Context) (interface{}, int32, error)
	GetCompany(key string, ctx context.Context) (interface{}, int32, error)
	GetUser(key string, ctx context.Context) (interface{}, int32, error)
	Set(key string, value interface{}, ctx context.Context) (int32, error)
	Delete(key string, ctx context.Context) (int32, error)
}

type Cacher struct {
	client  *redis.Client
	host    string
	port    string
	db      int
	expires time.Duration
}

func NewRedisCache(host string, port string, db int, expires time.Duration) *Cacher {
	return &Cacher{
		client:  redis.NewClient(&redis.Options{Addr: host + ":" + port}),
		host:    host,
		port:    port,
		db:      db,
		expires: expires,
	}
}
