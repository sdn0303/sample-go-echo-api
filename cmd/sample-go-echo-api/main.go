package main

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/sdn0303/sample-go-echo-api/config"
	"github.com/sdn0303/sample-go-echo-api/internal/app/infrastructure/adapters"
	"github.com/sdn0303/sample-go-echo-api/internal/app/interfaces/routes"
	"github.com/sdn0303/sample-go-echo-api/pkg/utils"
)

var (
	err        error
	conf       *config.Config
	pool       *adapters.Database
	errHandler utils.ErrorHandlerIFace
)

func init() {
	conf = config.New()
	errHandler = utils.NewErrorHandler()
	pool, err = adapters.NewDatabase(conf.DatabaseURI)
	errHandler.Fatalf(err, "failed connecting to mysql")
}

func createMiddleware(e *echo.Echo) {
	e.Use(middleware.Recover())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${time_rfc3339_nano} ${host} ${method} ${uri} ${status} \n",
	}))
}

func main() {
	e := echo.New()
	createMiddleware(e)

	routing := routes.New(e, conf, pool)
	routing.Bind(e)

	log.Fatal(e.Start(":8080").Error())
}
