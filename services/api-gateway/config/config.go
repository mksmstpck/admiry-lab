package config

import (
	"log"
	"os"
	"strconv"
	"time"
)

type Config struct {
	NatsUrl           string
	GinUrl            string
	TTLAccess         time.Duration
	TTLRefresh        time.Duration
	PrivateRefreshKey []byte
	PublicRefreshKey  []byte
	PrivateAccessKey  []byte
	PublicAccessKey   []byte
}

func NewConfig() Config {
	ttlAccess, err := strconv.Atoi(os.Getenv("TTL_ACCESS"))
	if err != nil {
		log.Fatal(err)
	}

	ttlRefresh, err := strconv.Atoi(os.Getenv("TTL_REFRESH"))
	if err != nil {
		log.Fatal(err)
	}

	return Config{
		NatsUrl:           os.Getenv("NATS_URI"),
		GinUrl:            os.Getenv("GIN_URL"),
		TTLAccess:         time.Duration(ttlAccess) * time.Second,
		TTLRefresh:        time.Duration(ttlRefresh) * time.Hour,
		PrivateRefreshKey: []byte(os.Getenv("REFRESH_SECRET")),
		PublicRefreshKey:  []byte(os.Getenv("REFRESH_PUBLIC")),
		PrivateAccessKey:  []byte(os.Getenv("ACCESS_SECRET")),
		PublicAccessKey:   []byte(os.Getenv("ACCESS_PUBLIC")),
	}
}
