package config

import "os"

type Config struct {
	NatsUrl string
	EchoUrl string
}

func NewConfig() Config {
	return Config{
		NatsUrl: os.Getenv("NATS_URI"),
		EchoUrl: os.Getenv("ECHO_URL"),
	}
}
