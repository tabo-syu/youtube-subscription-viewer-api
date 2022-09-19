package main

import (
	"log"

	"github.com/tabo-syu/youtube-subscription-viewer-api/infrastructures"
)

var (
	sql     *infrastructures.SqlHandler
	youtube *infrastructures.YoutubeHandler
	err     error
)

func init() {
	config := infrastructures.NewConfig()

	sql, err = infrastructures.NewSqlHandler(&config.DB)
	if err != nil {
		log.Fatalf("Cannot connect DB")
	}

	youtube = infrastructures.NewYoutubeHandler(&config.Youtube)
	// if err != nil {
	// 	log.Fatalf("Cannot connect YouTube API")
	// }
}

func main() {
	defer sql.Close()

	server, err := infrastructures.NewServer(sql, youtube)
	if err != nil {
		log.Fatalf("Cannot start the server")
	}

	server.Start("8080")
}
