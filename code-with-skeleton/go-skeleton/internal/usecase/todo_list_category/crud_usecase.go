package todo_list_usecase

import (
	"context"
	"fmt"
	"time"

	errwrap "github.com/pkg/errors"
	generalEntity "github.com/rahmatrdn/go-skeleton/entity"
	"github.com/rahmatrdn/go-skeleton/internal/helper"
	"github.com/rahmatrdn/go-skeleton/internal/repository/mysql"
	mentity "github.com/rahmatrdn/go-skeleton/internal/repository/mysql/entity"
	"github.com/rahmatrdn/go-skeleton/internal/usecase"
	"github.com/rahmatrdn/go-skeleton/internal/usecase/todo_list_category/entity"
)

type CrudTodoListCategoryUsecase struct {
	todoListCategoryRepo mysql.ITodoListCategoryRepository
}

func NewCrudTodoListCategoryUsecase(
	todoListCategoryRepo mysql.ITodoListCategoryRepository,
) *CrudTodoListCategoryUsecase {
	return &CrudTodoListCategoryUsecase{todoListCategoryRepo}
}

type ICrudTodoListCategoryUsecase interface {
	GetByID(ctx context.Context, todoListCategoryID int64) (*entity.TodoListCategoryResponse, error)
	GetAll(ctx context.Context) (res []*entity.TodoListCategoryResponse, err error)
	Create(ctx context.Context, TodoListCategoryReq entity.TodoListCategoryReq) (*entity.TodoListCategoryResponse, error)
	UpdateByID(ctx context.Context, TodoListCategoryReq entity.TodoListCategoryReq) error
	DeleteByID(ctx context.Context, todoListCategoryID int64) error
}

func (t *CrudTodoListCategoryUsecase) GetByID(ctx context.Context, todoListCategoryID int64) (*entity.TodoListCategoryResponse, error) {
	funcName := "CrudTodoListCategoryUsecase.GetByID"
	captureFieldError := generalEntity.CaptureFields{}

	data, err := t.todoListCategoryRepo.GetByID(ctx, todoListCategoryID)
	if err != nil {
		helper.LogError("todoListCategoryRepo.GetByID", funcName, err, captureFieldError, "")

		return nil, err
	}
	if data == nil {
		return nil, nil
	}

	return &entity.TodoListCategoryResponse{
		ID:          data.ID,
		Name:        data.Name,
		Description: data.Description,
		CreatedAt:   helper.ConvertToJakartaTime(data.CreatedAt),
	}, nil
}

func (t *CrudTodoListCategoryUsecase) GetAll(ctx context.Context) (res []*entity.TodoListCategoryResponse, err error) {
	funcName := "CrudTodoListCategoryUsecase.GetAll"
	captureFieldError := generalEntity.CaptureFields{}

	result, err := t.todoListCategoryRepo.GetAll(ctx)
	if err != nil {
		helper.LogError("todoListCategoryRepo.GetAll", funcName, err, captureFieldError, "")

		return nil, err
	}

	for _, v := range result {
		res = append(res, &entity.TodoListCategoryResponse{
			ID:          v.ID,
			Name:        v.Name,
			Description: v.Description,
			CreatedAt:   helper.ConvertToJakartaTime(v.CreatedAt),
		})
	}

	return res, nil
}

func (t *CrudTodoListCategoryUsecase) Create(ctx context.Context, TodoListCategoryReq entity.TodoListCategoryReq) (*entity.TodoListCategoryResponse, error) {
	funcName := "CrudTodoListCategoryUsecase.Create"
	captureFieldError := generalEntity.CaptureFields{
		"payload": helper.ToString(TodoListCategoryReq),
	}

	if errMsg := usecase.ValidateStruct(TodoListCategoryReq); errMsg != "" {
		return nil, errwrap.Wrap(fmt.Errorf(generalEntity.INVALID_PAYLOAD_CODE), errMsg)
	}

	todoListCategoryPayload := &mentity.TodoListCategory{
		Name:        TodoListCategoryReq.Name,
		Description: TodoListCategoryReq.Description,
		CreatedAt:   time.Now(),
	}

	err := t.todoListCategoryRepo.Create(ctx, nil, todoListCategoryPayload, false)
	if err != nil {
		helper.LogError("todoListCategoryRepo.Create", funcName, err, captureFieldError, "")

		return nil, err
	}

	return &entity.TodoListCategoryResponse{
		ID:          todoListCategoryPayload.ID,
		Name:        todoListCategoryPayload.Name,
		Description: todoListCategoryPayload.Description,
		CreatedAt:   helper.ConvertToJakartaTime(todoListCategoryPayload.CreatedAt),
	}, nil
}

func (t *CrudTodoListCategoryUsecase) UpdateByID(ctx context.Context, TodoListCategoryReq entity.TodoListCategoryReq) error {
	funcName := "CrudTodoListCategoryUsecase.UpdateByID"
	todoListCategoryID := TodoListCategoryReq.ID

	captureFieldError := generalEntity.CaptureFields{
		"payload": helper.ToString(TodoListCategoryReq),
	}

	if err := mysql.DBTransaction(t.todoListCategoryRepo, func(trx mysql.TrxObj) error {
		lockedData, err := t.todoListCategoryRepo.LockByID(ctx, trx, todoListCategoryID)
		if err != nil {
			helper.LogError("todoListCategoryRepo.LockByID", funcName, err, captureFieldError, "")

			return err
		}
		if lockedData == nil {
			return fmt.Errorf("DATA IS NOT EXIST")
		}

		if err := t.todoListCategoryRepo.Update(ctx, trx, lockedData, &mentity.TodoListCategory{
			Name:        TodoListCategoryReq.Name,
			Description: TodoListCategoryReq.Description,
		}); err != nil {
			helper.LogError("todoListCategoryRepo.Update", funcName, err, captureFieldError, "")

			return err
		}

		return nil
	}); err != nil {
		helper.LogError("todoListCategoryRepo.DBTransaction", funcName, err, captureFieldError, "")

		return err
	}

	return nil
}

func (t *CrudTodoListCategoryUsecase) DeleteByID(ctx context.Context, todoListCategoryID int64) error {
	funcName := "CrudTodoListCategoryUsecase.DeleteByID"
	captureFieldError := generalEntity.CaptureFields{
		"todo_list_id": helper.ToString(todoListCategoryID),
	}

	err := t.todoListCategoryRepo.DeleteByID(ctx, nil, todoListCategoryID)
	if err != nil {
		helper.LogError("todoListCategoryRepo.DeleteByID", funcName, err, captureFieldError, "")

		return err
	}

	return nil
}
