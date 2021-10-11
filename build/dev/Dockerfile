FROM golang:1.17-alpine
WORKDIR $GOPATH/src/github.com/sdn0303/smaple-go-echo-api

COPY ./go.mod .
COPY ./go.sum .
RUN set -ex
RUN apk update && apk add --no-cache gcc make git alpine-sdk build-base libc-dev binutils-gold
RUN go mod download && go get github.com/cosmtrek/air
RUN go get -d \
      golang.org/x/tools/cmd/goimports@latest \
      github.com/golangci/golangci-lint/cmd/golangci-lint@v1.42.0 \
      entgo.io/ent/cmd/ent \
      golang.org/x/lint/golint \
      github.com/swaggo/swag/cmd/swag \
      github.com/golang/mock/mockgen \
      github.com/fzipp/gocyclo/cmd/gocyclo@latest \
      github.com/BurntSushi/toml/cmd/tomlv@master \
      honnef.co/go/tools/cmd/staticcheck@latest

COPY . .

EXPOSE 8080
