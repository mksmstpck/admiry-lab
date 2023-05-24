package config

import (
	"os"
	"strconv"
)

type Config struct {
	PSQLhost     string
	PSQLport     string
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
		panic(err)
	}
	redisExpiresint, err := strconv.Atoi(os.Getenv("REDIS_EXPIRES"))
	if err != nil {
		panic(err)
	}

	return Config{
		PSQLhost:     os.Getenv("PSQL_HOST"),
		PSQLport:     os.Getenv("PSQL_PORT"),
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
