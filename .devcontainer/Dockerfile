FROM golang:1.19-bullseye

RUN apt-get update && \
    # install golang devtools
    go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest && \
    go install golang.org/x/tools/gopls@latest && \
    go install github.com/ramya-rao-a/go-outline@latest && \
    go install github.com/go-delve/delve/cmd/dlv@latest && \
    go install honnef.co/go/tools/cmd/staticcheck@latest

WORKDIR /go/src/youtube-channel-manager-api
