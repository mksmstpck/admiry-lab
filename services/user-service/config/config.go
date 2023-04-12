package config

import "github.com/spf13/viper"

type Config struct {
	PSQLhost     string `mapstructure:"PSQL_HOST"`
	PSQLport     string `mapstructure:"PSQL_PORT"`
	PSQLuser     string `mapstructure:"PSQL_USER"`
	PSQLpass     string `mapstructure:"PSQL_PASS"`
	PSQLdb       string `mapstructure:"PSQL_DB"`
	NatsUrl      string `mapstructure:"NATS_URL"`
	RedisHost    string `mapstructure:"REDIS_HOST"`
	RedisPort    string `mapstructure:"REDIS_PORT"`
	RedisDB      int    `mapstructure:"REDIS_DB"`
	RedisExpires int64  `mapstructure:"REDIS_EXPIRES"`
}

func NewConfig() (c Config, err error) {
	viper.AddConfigPath(".")
	viper.SetConfigName("dev")
	viper.SetConfigType("env")

	if err = viper.ReadInConfig(); err != nil {
		return
	}

	err = viper.Unmarshal(&c)

	return
}
