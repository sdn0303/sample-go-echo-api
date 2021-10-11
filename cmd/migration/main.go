package main

import (
	"context"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/sdn0303/sample-go-echo-api/pkg/ent"
	"github.com/sdn0303/sample-go-echo-api/pkg/ent/migrate"
	"github.com/sdn0303/sample-go-echo-api/pkg/utils"
)

var errHandler utils.ErrorHandlerIFace

func init() {
	errHandler = utils.NewErrorHandler()
}

func main() {

	var (
		client *ent.Client
		err    error
	)

	dbURI := os.Getenv("DATABASE_URI")
	client, err = ent.Open("mysql", dbURI)
	errHandler.Fatalf(err, "failed connecting to mysql")

	defer func(client *ent.Client) {
		err = client.Close()
		errHandler.Fatalf(err, "failed to close database connection")
	}(client)

	ctx := context.Background()
	err = client.Debug().Schema.Create(ctx, migrate.WithDropIndex(true), migrate.WithDropColumn(true))
	errHandler.Fatalf(err, "failed to migrate database schema")
}
