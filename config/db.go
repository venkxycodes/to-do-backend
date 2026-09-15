package config

import (
	"fmt"
	"net/url"
)

type DbConfig struct {
	Host     string `yaml:"HOST" env:"DB_HOST"`
	Username string `yaml:"USERNAME" env:"DB_USERNAME"`
	Password string `yaml:"PASSWORD" env:"DB_PASSWORD"`
	Port     string `yaml:"PORT" env:"DB_PORT"`
	DBName   string `yaml:"DB_NAME" env:"DB_NAME"`
}

type RedisConfig struct {
	RedisAddress string `yaml:"REDIS_ADDRESS" env:"REDIS_ADDRESS"`
}

func (dc *DbConfig) GetConnectionString() string {
	connectionString := fmt.Sprintf("mongodb://%s:%s", dc.Host, dc.Port)
	if dc.Username != "" {
		credentials := url.UserPassword(dc.Username, dc.Password).String()
		connectionString = fmt.Sprintf("mongodb://%s@%s:%s", credentials, dc.Host, dc.Port)
	}
	return connectionString
}
