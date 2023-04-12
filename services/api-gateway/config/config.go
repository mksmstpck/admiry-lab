package config

import "github.com/spf13/viper"

type Config struct {
	NatsUrl string `mapstructure:"NATS_URL"`
	EchoUrl string `mapstructure:"ECHO_URL"`
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
