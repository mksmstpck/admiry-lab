package conectors

import (
	"time"

	"github.com/labstack/gommon/log"
	"github.com/mkskstpck/admiry-lab/pkg/cache"
)

func NewCache(host string, port string, db int, expires time.Duration) *cache.Cacher {
	log.Info("conect to redis db")
	return cache.NewRedisCache(
		host,
		port,
		db,
		expires,
	)
}
