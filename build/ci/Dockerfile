FROM golang:1.17-alpine
ENV GO111MODULE=on

RUN apk update && apk add --no-cache gcc make git alpine-sdk build-base libc-dev binutils-gold
RUN apt-get -y install jq \
RUN go get -d \
		github.com/kisielk/errcheck \
		golang.org/x/tools/cmd/goimports \
		golang.org/x/lint/golint \
		github.com/securego/gosec/cmd/gosec \
		golang.org/x/tools/go/analysis/passes/shadow/cmd/shadow \
		honnef.co/go/tools/cmd/staticcheck

COPY ./scripts/entrypoint.sh /entrypoint.sh
ENTRYPOINT ["/entrypoint.sh"]