package config

import (
	"strings"

	"github.com/spf13/viper"
)

type Postgres struct {
	Host     string
	User     string
	Password string
}

type Config struct {
	Postgres *Postgres
}

func Read() (*Config, error) {
	viper.AutomaticEnv()
	viper.SetEnvPrefix("APP")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))

	viper.SetConfigName("microservice")
	viper.SetConfigType("yml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("$HOME")
	if err := viper.ReadInConfig(); err != nil { // Handle errors reading the config file
		return nil, err
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil { // Handle errors reading the config file
		return nil, err
	}
	return &config, nil
}
