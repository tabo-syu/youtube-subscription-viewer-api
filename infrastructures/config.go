package infrastructures

import (
	"os"
)

type DB struct {
	Host     string
	Port     string
	Name     string
	User     string
	Password string
	TimeZone string
}

type Youtube struct {
	ClientSecret []byte
}

type Config struct {
	DB      DB
	Youtube Youtube
}

func NewConfig() (*Config, error) {
	secret, err := os.ReadFile("client_secret.json")
	if err != nil {
		return nil, err
	}

	return &Config{
		DB: DB{
			os.Getenv("DB_HOST"),
			os.Getenv("DB_PORT"),
			os.Getenv("DB_NAME"),
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_TIMEZONE"),
		},
		Youtube: Youtube{secret},
	}, nil
}
