package infrastructures

import "os"

type DB struct {
	Host     string
	Port     string
	Name     string
	User     string
	Password string
	TimeZone string
}

type Youtube struct{}

type Config struct {
	DB      DB
	Youtube Youtube
}

func NewConfig() *Config {
	return &Config{
		DB: DB{
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
			Name:     os.Getenv("DB_NAME"),
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			TimeZone: os.Getenv("DB_TIMEZONE"),
		},
		Youtube: Youtube{},
	}
}
