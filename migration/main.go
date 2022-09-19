package migration

import (
	"fmt"
	"log"

	"github.com/tabo-syu/youtube-subscription-viewer-api/infrastructures"
)

func main() {
	sql, err := infrastructures.NewSqlHandler(
		infrastructures.NewConfig().DB,
	)
	if err != nil {
		log.Fatalf("Could not connect DB")
	}
	defer sql.Close()

	if err := migrate(); err != nil {
		log.Fatalf("Could not migrate")
	}

	fmt.Println("Migration is complete!")
}

func migrate() error {
	return nil
}
