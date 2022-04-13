package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Postgres Postgres
	Rabbit   Rabbit
	Server   Server
}

type Postgres struct {
	Host          string
	Port          string
	Username      string
	Password      string
	Database      string
	Debug         bool
	SslMode       string `mapstructure:"ssl-mode"`
	MigrationPath string `mapstructure:"migration-path"`
}

type Rabbit struct {
	Username string
	Password string
	Host     string
}

type Server struct {
	Port int
}

func Load(path string) (*Config, error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app-config")
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}
	var conf Config
	if err := viper.Unmarshal(&conf); err != nil {
		return nil, err
	}
	return &conf, nil
}
