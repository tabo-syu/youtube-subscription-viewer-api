package main

import (
	"github.com/tabo-syu/youtube-subscription-viewer-api/infrastructure"
)

func main() {
	infrastructure.Serve("8080")
}
