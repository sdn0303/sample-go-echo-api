package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/sdn0303/sample-go-echo-api/internal/app/domain/model"
	"github.com/sdn0303/sample-go-echo-api/internal/app/usecase"
	"github.com/sdn0303/sample-go-echo-api/pkg/utils"
)

// TodoHandlerIFace handlers interface
type TodoHandlerIFace interface {
	ListTodo() echo.HandlerFunc
	GetTodo() echo.HandlerFunc
	CreateTodo() echo.HandlerFunc
	PatchTodo() echo.HandlerFunc
	DeleteTodo() echo.HandlerFunc
}

type handler struct {
	Usecase usecase.TodoUsecaseIFace
}

// NewTodoHandler constructs a new handler
func NewTodoHandler(usecase usecase.TodoUsecaseIFace) TodoHandlerIFace {
	return &handler{Usecase: usecase}
}

// ListTodo get all todo
func (th *handler) ListTodo() echo.HandlerFunc {
	return func(c echo.Context) error {

		ctx := context.Background()
		listTodo, err := th.Usecase.ListTodo(ctx)
		if err != nil {
			log.Println("failed to get list todo")
			return echo.NewHTTPError(http.StatusInternalServerError, err)
		}

		return c.JSON(http.StatusOK, listTodo)
	}
}

// GetTodo get todo
func (th *handler) GetTodo() echo.HandlerFunc {
	return func(c echo.Context) error {

		ctx := context.Background()
		idStr := c.Param("id")
		id, err := utils.StringToInt(idStr)
		if err != nil {
			log.Printf("failed to convert string to int: %v\n", err)
			return echo.NewHTTPError(http.StatusInternalServerError, err)
		}

		todo, err := th.Usecase.GetTodo(ctx, id)
		if err != nil {
			log.Printf("failed to get todo: %v\n", err)
			return echo.NewHTTPError(http.StatusInternalServerError, err)
		}

		return c.JSON(http.StatusOK, todo)
	}
}

// CreateTodo create todo
func (th *handler) CreateTodo() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := context.Background()

		t := new(model.TodoModel)
		if err := c.Bind(&t); err != nil {
			log.Printf("failed to bind params: %v\n", err)
			return echo.NewHTTPError(http.StatusInternalServerError, err)
		}

		log.Println(fmt.Printf("todoModel=%v", t))

		if err := th.Usecase.CreateTodo(ctx, t); err != nil {
			log.Printf("failed to create todo: %v\n", err)
			return echo.NewHTTPError(http.StatusInternalServerError, err)
		}

		return c.JSON(http.StatusOK, "todo created")
	}
}

// PatchTodo patch todo
func (th *handler) PatchTodo() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := context.Background()
		idStr := c.Param("id")

		id, err := utils.StringToInt(idStr)
		if err != nil {
			log.Printf("failed to convert string to int: %v\n", err)
			return echo.NewHTTPError(http.StatusInternalServerError, err)
		}

		t := new(model.TodoModel)
		if err := c.Bind(&t); err != nil {
			log.Printf("failed to bind params: %v\n", err)
			return echo.NewHTTPError(http.StatusInternalServerError, err)
		}

		log.Println(fmt.Printf("id=%v", id))
		log.Println(fmt.Printf("todoModel=%v", t))

		if err := th.Usecase.PatchTodo(ctx, id, t); err != nil {
			log.Printf("failed to update todo: %v\n", err)
			return echo.NewHTTPError(http.StatusInternalServerError, err)
		}

		return c.JSON(http.StatusOK, "todo updated")
	}
}

// DeleteTodo delete todo
func (th *handler) DeleteTodo() echo.HandlerFunc {
	return func(c echo.Context) error {

		ctx := context.Background()
		idStr := c.Param("id")

		id, err := utils.StringToInt(idStr)
		if err != nil {
			log.Printf("failed to convert string to int: %v\n", err)
			return echo.NewHTTPError(http.StatusInternalServerError, err)
		}

		log.Println(fmt.Printf("id=%v", id))

		if err := th.Usecase.DeleteTodo(ctx, id); err != nil {
			log.Printf("failed to deleted todo: %v\n", err)
			return echo.NewHTTPError(http.StatusInternalServerError, err)
		}

		return c.JSON(http.StatusOK, "todo deleted")
	}
}
