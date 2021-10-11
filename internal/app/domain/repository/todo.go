package repository

import (
	"context"

	"github.com/sdn0303/sample-go-echo-api/internal/app/domain/model"
)

// TodoRepositoryIFace interface of repository
type TodoRepositoryIFace interface {
	All(ctx context.Context) ([]*model.TodoModel, error)
	FindByID(ctx context.Context, id int) (*model.TodoModel, error)
	Create(ctx context.Context, todo *model.TodoModel) error
	Update(ctx context.Context, id int, todo *model.TodoModel) error
	Delete(ctx context.Context, is int) error
}
