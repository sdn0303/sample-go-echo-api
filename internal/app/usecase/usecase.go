package usecase

import (
	"context"

	"github.com/sdn0303/sample-go-echo-api/internal/app/domain/model"
	"github.com/sdn0303/sample-go-echo-api/internal/app/domain/repository"
)

// TodoUsecaseIFace usecase interface
type TodoUsecaseIFace interface {
	ListTodo(ctx context.Context) ([]*model.TodoModel, error)
	GetTodo(ctx context.Context, id int) (*model.TodoModel, error)
	CreateTodo(ctx context.Context, todoModel *model.TodoModel) error
	PatchTodo(ctx context.Context, id int, todoModel *model.TodoModel) error
	DeleteTodo(ctx context.Context, id int) error
}

type usecase struct {
	Repository repository.TodoRepositoryIFace
}

// NewTodoUsecase constructs a new usecase
func NewTodoUsecase(repo repository.TodoRepositoryIFace) TodoUsecaseIFace {
	return &usecase{Repository: repo}
}

// ListTodo get all todo
func (tu *usecase) ListTodo(ctx context.Context) ([]*model.TodoModel, error) {
	return tu.Repository.All(ctx)
}

// GetTodo get todo
func (tu *usecase) GetTodo(ctx context.Context, id int) (*model.TodoModel, error) {
	return tu.Repository.FindByID(ctx, id)
}

// CreateTodo create todo
func (tu *usecase) CreateTodo(ctx context.Context, todoModel *model.TodoModel) error {
	return tu.Repository.Create(ctx, todoModel)
}

// PatchTodo patch todo
func (tu *usecase) PatchTodo(ctx context.Context, id int, todoModel *model.TodoModel) error {
	return tu.Repository.Update(ctx, id, todoModel)
}

// DeleteTodo delete todo
func (tu *usecase) DeleteTodo(ctx context.Context, id int) error {
	return tu.Repository.Delete(ctx, id)
}
