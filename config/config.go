package config

import (
	"fmt"
	"log/slog"
	"os"
)

type Config struct {
	Database database
	Redis    redis
	Http     http
	Email    email

	SecretKey     string
	TelegramToken string
}

func New() Config {
	return Config{
		Database: newDatabase(),
		Redis:    newRedis(),
		Http:     newHTTP(),

		SecretKey:     getEnv("JWT_SECRET_KEY"),
		TelegramToken: getEnv("TELEGRAM_TOKEN"),
	}
}

type database struct {
	name     string
	user     string
	password string
	host     string
	port     string
}

func newDatabase() database {
	return database{
		name:     getEnv("DATABASE_NAME"),
		user:     getEnv("DATABASE_USER"),
		password: getEnv("DATABASE_PASSWORD"),
		host:     getEnv("DATABASE_HOST"),
		port:     getEnv("DATABASE_PORT"),
	}
}

func (d database) Uri() string {
	return fmt.Sprintf("%s:%s@(%s:%s)/%s?parseTime=true", d.user, d.password, d.host, d.port, d.name)
}

type redis struct {
	host     string
	port     string
	password string
}

func newRedis() redis {
	return redis{
		host:     getEnv("REDIS_HOST"),
		port:     getEnv("REDIS_PORT"),
		password: getEnv("REDIS_PASSWORD"),
	}
}

func (r redis) Addr() string {
	return r.host + ":" + r.port
}

func (r redis) Password() string {
	return r.password
}

type http struct {
	host string
	port string
}

func newHTTP() http {
	return http{
		port: getEnv("HTTP_PORT"),
	}
}

func (h http) Addr() string {
	return fmt.Sprintf(":%s", h.port)
}

type email struct {
	address  string
	password string
}

func newEmail() email {
	return email{
		address:  getEnv("MAIL_ADDRESS"),
		password: getEnv("MAIL_PASSWORD"),
	}
}

func getEnv(key string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	slog.Error(fmt.Sprintf("Env variable not found: %s", key))
	os.Exit(0)
	return ""
}
