package routes

import (
	"net/http"

	"github.com/sdn0303/sample-go-echo-api/internal/app/interfaces/handlers"
	"github.com/sdn0303/sample-go-echo-api/internal/app/usecase"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"

	"github.com/sdn0303/sample-go-echo-api/config"
	"github.com/sdn0303/sample-go-echo-api/internal/app/infrastructure/adapters"
	"github.com/sdn0303/sample-go-echo-api/internal/app/infrastructure/repository"
)

// Routing struct for routing
type Routing struct {
	e    *echo.Echo
	conf *config.Config
	pool *adapters.Database
}

// New constructs a new routing
func New(e *echo.Echo, conf *config.Config, pool *adapters.Database) *Routing {
	return &Routing{e: e, conf: conf, pool: pool}
}

func (r *Routing) todoRoutingGroup(e *echo.Echo) {
	// generate layers
	repo := repository.NewTodoRepository(r.pool)
	u := usecase.NewTodoUsecase(repo)
	h := handlers.NewTodoHandler(u)

	// generate routing group
	todo := e.Group("todo")
	todo.GET("/", h.ListTodo())
	todo.GET("/:id", h.GetTodo())
	todo.POST("/", h.CreateTodo())
	todo.PATCH("/:id", h.PatchTodo())
	todo.DELETE("/:id", h.DeleteTodo())
}

// Bind a routes
func (r *Routing) Bind(e *echo.Echo) {
	// Base Endpoints
	e.GET("/", func(c echo.Context) error { return c.String(http.StatusOK, "Hello, this is sample todo api server.") })
	e.GET("/docs/*", echoSwagger.WrapHandler)

	// todo endpoints
	r.todoRoutingGroup(e)
}
