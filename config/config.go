package config

import (
	"os"

	"github.com/sirupsen/logrus"

	"github.com/joho/godotenv"
)

type Config struct {
	DBUser           string
	DBPass           string
	DBName           string
	DBHost           string
	DBPort           string
	HttpPort         string
	CookieSecret     string
	CreateUserSecret string
	FTP              FTP
}

type DB struct {
	DBUser string
	DBPass string
	DBName string
	DBHost string
	DBPort string
}

type FTP struct {
	FilesPath string
}

func NewConfig() *Config {
	err := godotenv.Load(".env")
	if err != nil {
		logrus.Warn("no .env file, reading config from OS ENV variables")
	}
	return &Config{
		DBUser:           os.Getenv("DB_USER"),
		DBPass:           os.Getenv("DB_PASSWORD"),
		DBName:           os.Getenv("DB_NAME"),
		DBHost:           os.Getenv("DB_HOST"),
		DBPort:           os.Getenv("DB_PORT"),
		HttpPort:         os.Getenv("HTTP_PORT"),
		CookieSecret:     os.Getenv("COOKIE_SECRET"),
		CreateUserSecret: os.Getenv("CREATE_USER_SECRET"),
		FTP:              FTP{FilesPath: os.Getenv("WAVS_PATH")},
	}
}
