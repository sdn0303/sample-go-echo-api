package repository

import (
	"context"
	"fmt"
	"log"

	"github.com/sdn0303/sample-go-echo-api/internal/app/domain/model"
	"github.com/sdn0303/sample-go-echo-api/internal/app/domain/repository"
	"github.com/sdn0303/sample-go-echo-api/internal/app/infrastructure/adapters"
	"github.com/sdn0303/sample-go-echo-api/pkg/ent"
	"github.com/sdn0303/sample-go-echo-api/pkg/ent/todo"
)

// TodoRepository todo repository
type TodoRepository struct {
	Pool *adapters.Database
}

// NewTodoRepository constructs a new TodoRepository
func NewTodoRepository(pool *adapters.Database) repository.TodoRepositoryIFace {
	return &TodoRepository{
		Pool: pool,
	}
}

func dbObjToModel(dbObj *ent.Todo) *model.TodoModel {
	return &model.TodoModel{
		ID:          dbObj.ID,
		Title:       dbObj.Title,
		Description: dbObj.Description,
		Status:      dbObj.Status.String(),
	}
}

// All get all todo
func (tr *TodoRepository) All(ctx context.Context) ([]*model.TodoModel, error) {
	todoObjs, err := tr.Pool.Client.Todo.Query().All(ctx)
	if err != nil {
		return nil, err
	}

	var todoModels []*model.TodoModel
	for _, todoObj := range todoObjs {
		todoModels = append(todoModels, dbObjToModel(todoObj))
	}
	return todoModels, nil
}

// FindByID find todo by id
func (tr *TodoRepository) FindByID(ctx context.Context, id int) (*model.TodoModel, error) {
	todoObj, err := tr.Pool.Client.Todo.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	todoModel := dbObjToModel(todoObj)
	return todoModel, nil
}

// Create create todo
func (tr *TodoRepository) Create(ctx context.Context, todoModel *model.TodoModel) error {
	todoObj, err := tr.Pool.Client.Todo.Create().
		SetTitle(todoModel.Title).
		SetDescription(todoModel.Description).
		SetStatus(todo.Status(todoModel.Status)).
		Save(ctx)
	if err != nil {
		return err
	}
	log.Println(fmt.Sprintf("todoModel created: %v", todoObj))
	return nil
}

// Update update todo
func (tr *TodoRepository) Update(ctx context.Context, id int, todoModel *model.TodoModel) error {
	todoObj, err := tr.Pool.Client.Todo.UpdateOneID(id).
		SetTitle(todoModel.Title).
		SetDescription(todoModel.Description).
		SetStatus(todo.Status(todoModel.Status)).
		Save(ctx)
	if err != nil {
		return err
	}
	log.Println(fmt.Sprintf("todoModel updated: %v", todoObj))
	return nil
}

// Delete delete todo
func (tr *TodoRepository) Delete(ctx context.Context, id int) error {
	return tr.Pool.Client.Todo.DeleteOneID(id).Exec(ctx)
}
