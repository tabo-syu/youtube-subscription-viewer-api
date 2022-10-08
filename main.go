package main

import (
	"log"

	"github.com/tabo-syu/youtube-subscription-viewer-api/infrastructures"
	"github.com/tabo-syu/youtube-subscription-viewer-api/migration"
)

func main() {
	config, err := infrastructures.NewConfig()
	if err != nil {
		log.Fatalf("Cannot read client_secret.json")
	}

	sql, err := infrastructures.NewSQLHandler(&config.DB)
	if err != nil {
		log.Fatalf("Cannot connect DB")
	}
	defer sql.Close()

	if err := migration.Migrate(sql); err != nil {
		log.Fatalf("Cannot migrate cause:\n%s", err)
	}

	oAuth, err := infrastructures.NewYoutubeOAuth2Handler(config.Youtube.ClientSecret)
	if err != nil {
		log.Fatalf("Cannot initialize OAuth handler")
	}

	youtube := infrastructures.NewYoutubeHandler()

	server, err := infrastructures.NewServer(sql, oAuth, youtube)
	if err != nil {
		log.Fatalf("Cannot start the server")
	}

	server.Start("8080")
}
