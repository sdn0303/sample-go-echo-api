#!/bin/bash
set -ex

# setup Local environments
# set golang environments
export GO111MODULE=on
export GOPATH="${HOME}/go"
GO_VERSION=$(go version | grep -E -o "([0-9].[0-9]+.[0-9]*|\_[0-9])")
export GOROOT="/opt/homebrew/Cellar/go/${GO_VERSION}/libexec"
if [[ $PATH != *$GOPATH* ]]; then export PATH="${GOPATH}/bin:${PATH}"; fi
if [[ $PATH != *$GOROOT* ]]; then export PATH="${GOROOT}/bin:${PATH}"; fi

# install dev dependencies
go get github.com/cosmtrek/air
go get -d \
      golang.org/x/tools/cmd/goimports@latest \
      github.com/golangci/golangci-lint/cmd/golangci-lint@v1.42.0 \
      entgo.io/ent/cmd/ent \
      golang.org/x/lint/golint \
      github.com/swaggo/swag/cmd/swag \
      github.com/golang/mock/mockgen \
      github.com/fzipp/gocyclo/cmd/gocyclo@latest \
      github.com/BurntSushi/toml/cmd/tomlv@master \
      honnef.co/go/tools/cmd/staticcheck@latest
