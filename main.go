package main

import (
	"os"

	"github.com/tabo-syu/youtube-subscription-viewer-api/infrastructures"
)

var conf *infrastructures.Config

func init() {
	conf = &infrastructures.Config{
		DB: infrastructures.DB{
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
			Name:     os.Getenv("DB_NAME"),
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			TimeZone: os.Getenv("DB_TIMEZONE"),
		},
		Youtube: infrastructures.Youtube{},
	}
}

func main() {
	server := infrastructures.NewServer(conf)
	server.Start("8080")
}
