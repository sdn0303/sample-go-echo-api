FROM golang:1.17-alpine as builder
WORKDIR $GOPATH/src/github.com/sdn0303/smaple-go-echo-api

COPY ./go.mod .
COPY ./go.sum .
RUN set -ex
RUN apk update && apk add --no-cache gcc make git alpine-sdk build-base libc-dev binutils-gold
RUN go mod download

COPY ./cmd ./cmd
COPY ./config ./config
COPY ./internal ./internal
RUN CGO_ENABLED=0 GOOS=linux go build -i -v -o ./bin/sample-go-echo-api ./cmd/sample-go-echo-api/main.go


FROM golang:1.16-alpine as production
WORKDIR /usr/local/bin
COPY --from=builder /go/src/github.com/sdn0303/bin/smaple-go-echo-api /smaple-go-echo-api
ENTRYPOINT ["/smaple-go-echo-api"]