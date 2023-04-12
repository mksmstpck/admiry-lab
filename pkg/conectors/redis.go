package conectors

import (
	"time"

	"github.com/mkskstpck/to-rename/pkg/cache"
)

func NewCache(host string, port string, db int, expires time.Duration) *cache.Cacher {
	return cache.NewRedisCache(
		host,
		port,
		db,
		expires,
	)
}
