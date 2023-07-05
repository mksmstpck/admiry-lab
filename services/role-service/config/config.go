package config

import (
	"os"
	"strconv"

	"github.com/labstack/gommon/log"
)

type Config struct {
	PSQLaddr     string
	PSQLuser     string
	PSQLpass     string
	PSQLdb       string
	NatsURI      string
	RedisHost    string
	RedisPort    string
	RedisDB      int
	RedisExpires int
}

func NewConfig() Config {
	redisDBint, err := strconv.Atoi(os.Getenv("REDIS_DB"))
	if err != nil {
		log.Fatal(err)
	}
	redisExpiresint, err := strconv.Atoi(os.Getenv("REDIS_EXPIRES"))
	if err != nil {
		log.Fatal(err)
	}

	return Config{
		PSQLaddr:     os.Getenv("PSQL_ADDR"),
		PSQLuser:     os.Getenv("PSQL_USER"),
		PSQLpass:     os.Getenv("PSQL_PASS"),
		PSQLdb:       os.Getenv("PSQL_DB"),
		NatsURI:      os.Getenv("NATS_URI"),
		RedisHost:    os.Getenv("REDIS_HOST"),
		RedisPort:    os.Getenv("REDIS_PORT"),
		RedisDB:      redisDBint,
		RedisExpires: redisExpiresint,
	}
}
