package config

import (
	"os"
)

type Config struct {
	AppName     string   `yaml:"APP_NAME" env:"APP_NAME"`
	AppPort     string   `yaml:"APP_PORT" env:"APP_PORT"`
	ENV         string   `yaml:"ENV" env:"ENVIRONMENT"`
	JWTSecret   string   `yaml:"JWT_SECRET" env:"JWT_SECRET"`
	CORSOrigins string   `yaml:"CORS_ORIGINS" env:"CORS_ORIGINS"`
	DbConfig    DbConfig `yaml:"DB_CONFIG" env:"DB_CONFIG"`
}

func (c *Config) SetDefault() {
	c.AppName = "to-do"
	c.AppPort = "9090"
	c.ENV = "dev"
	c.JWTSecret = ""
	c.CORSOrigins = ""
	c.DbConfig = DbConfig{
		Host:     "localhost",
		Username: "",
		Password: "",
		Port:     "27017",
		DBName:   "to-do",
	}
}

func GetConfig() *Config {
	config := &Config{}
	config.SetDefault()
	if value := os.Getenv("ENV"); value != "" {
		config.ENV = value
	}
	if value := os.Getenv("APP_NAME"); value != "" {
		config.AppName = value
	}
	if value := os.Getenv("APP_PORT"); value != "" {
		config.AppPort = value
	}
	if value := os.Getenv("JWT_SECRET"); value != "" {
		config.JWTSecret = value
	}
	if value := os.Getenv("CORS_ORIGINS"); value != "" {
		config.CORSOrigins = value
	}
	config.DbConfig = GetDbConfig()
	return config
}

func GetDbConfig() DbConfig {
	config := DbConfig{Host: "localhost", Port: "27017", DBName: "to-do"}
	if value := os.Getenv("DB_HOST"); value != "" {
		config.Host = value
	}
	if value := os.Getenv("DB_USERNAME"); value != "" {
		config.Username = value
	}
	if value := os.Getenv("DB_PASSWORD"); value != "" {
		config.Password = value
	}
	if value := os.Getenv("DB_PORT"); value != "" {
		config.Port = value
	}
	if value := os.Getenv("DB_NAME"); value != "" {
		config.DBName = value
	}
	return config
}
