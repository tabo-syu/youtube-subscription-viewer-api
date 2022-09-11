# TODO: マルチステージビルド
FROM golang:1.19-bullseye

WORKDIR /go/src/youtube-channel-manager-api

CMD [ "go", "run", "main.go" ]
