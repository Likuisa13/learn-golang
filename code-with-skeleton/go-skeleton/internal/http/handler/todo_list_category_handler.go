package handler

import (
	"net/http"

	"github.com/rahmatrdn/go-skeleton/internal/parser"
	"github.com/rahmatrdn/go-skeleton/internal/presenter/json"
	todo_list_category_usecase "github.com/rahmatrdn/go-skeleton/internal/usecase/todo_list_category"
	"github.com/rahmatrdn/go-skeleton/internal/usecase/todo_list_category/entity"

	fiber "github.com/gofiber/fiber/v2"
)

type TodoListCategoryHandler struct {
	parser                      parser.Parser
	presenter                   json.JsonPresenter
	todoListCategoryCrudUsecase todo_list_category_usecase.ICrudTodoListCategoryUsecase
}

func NewTodoListCategoryHandler(
	parser parser.Parser,
	presenter json.JsonPresenter,
	todoListCategoryCrudUsecase todo_list_category_usecase.ICrudTodoListCategoryUsecase,
) *TodoListCategoryHandler {
	return &TodoListCategoryHandler{parser, presenter, todoListCategoryCrudUsecase}
}

func (w *TodoListCategoryHandler) Register(app fiber.Router) {
	app.Get("/todo-list-category/:id", w.GetByID)
	app.Get("/todo-list-category", w.GetAll)
	app.Post("/todo-list-category", w.Create)
	app.Put("/todo-list-category/:id", w.Update)
	app.Delete("/todo-list-category/:id", w.Delete)
}

// @Summary         Get Todo List by ID
// @Description     Get a Todo List by its ID
// @Tags            Todo List
// @Accept          json
// @Produce         json
// @Security        Bearer
// @Param           id path int true "ID of the Todo List"
// @Success			201 {object} entity.GeneralResponse{data=entity.TodoListResponse} "Success"
// @Failure			401 {object} entity.CustomErrorResponse "Unauthorized"
// @Failure			422 {object} entity.CustomErrorResponse "Invalid Request Body"
// @Failure			500 {object} entity.CustomErrorResponse "Internal server Error"
// @Router			/api/v1/todo-lists/{id} [get]
func (w *TodoListCategoryHandler) GetByID(c *fiber.Ctx) error {
	id, err := w.parser.ParserIntIDFromPathParams(c)
	if err != nil {
		return w.presenter.BuildError(c, err)
	}

	data, err := w.todoListCategoryCrudUsecase.GetByID(c.Context(), id)
	if err != nil {
		return w.presenter.BuildError(c, err)
	}

	return w.presenter.BuildSuccess(c, data, "Success", http.StatusOK)
}

// @Summary         Retrieve Todo Lists by User ID
// @Description     Retrieve a list of Todo Lists belonging to a user by their User ID
// @Tags            Todo List
// @Accept			json
// @Produce			json
// @Security 		Bearer
// @Success			201 {object} entity.GeneralResponse{data=[]entity.TodoListResponse} "Success"
// @Failure			401 {object} entity.CustomErrorResponse "Unauthorized"
// @Failure			422 {object} entity.CustomErrorResponse "Invalid Request Body"
// @Failure			500 {object} entity.CustomErrorResponse "Internal server Error"
// @Router			/api/v1/todo-list [get]
func (w *TodoListCategoryHandler) GetAll(c *fiber.Ctx) error {
	data, err := w.todoListCategoryCrudUsecase.GetAll(c.Context())
	if err != nil {
		return w.presenter.BuildError(c, err)
	}

	return w.presenter.BuildSuccess(c, data, "Success", http.StatusOK)
}

// @Summary			Create a new Todo List
// @Description		Create a new Todo List
// @Tags			Todo List
// @Accept			json
// @Produce			json
// @Security 		Bearer
// @Param			req body entity.TodoListCategoryReq true "Payload Request Body"
// @Success			201 {object} entity.GeneralResponse{data=entity.TodoListCategoryReq} "Success"
// @Failure			401 {object} entity.CustomErrorResponse "Unauthorized"
// @Failure			422 {object} entity.CustomErrorResponse "Invalid Request Body"
// @Failure			500 {object} entity.CustomErrorResponse "Internal server Error"
// @Router			/api/v1/todo-list-category [post]
func (w *TodoListCategoryHandler) Create(c *fiber.Ctx) error {
	var req entity.TodoListCategoryReq

	err := w.parser.ParserBodyRequest(c, &req)
	if err != nil {
		return w.presenter.BuildError(c, err)
	}

	data, err := w.todoListCategoryCrudUsecase.Create(c.Context(), req)
	if err != nil {
		return w.presenter.BuildError(c, err)
	}

	return w.presenter.BuildSuccess(c, data, "Success", http.StatusOK)
}

// @Summary         Update an existing Todo List by ID
// @Description     Update an existing Todo List
// @Tags            Todo List
// @Accept          json
// @Produce         json
// @Security        Bearer
// @Param           id path int true "ID of the todo list"
// @Param			req body entity.TodoListCategoryReq true "Payload Request Body"
// @Success			201 {object} entity.GeneralResponse "Success"
// @Failure			401 {object} entity.CustomErrorResponse "Unauthorized"
// @Failure			422 {object} entity.CustomErrorResponse "Invalid Request Body"
// @Failure			500 {object} entity.CustomErrorResponse "Internal server Error"
// @Router			/api/v1/todo-list-category [put]
func (w *TodoListCategoryHandler) Update(c *fiber.Ctx) error {
	var req entity.TodoListCategoryReq
	err := w.parser.ParserBodyWithIntIDPathParams(c, &req)
	if err != nil {
		return w.presenter.BuildError(c, err)
	}

	err = w.todoListCategoryCrudUsecase.UpdateByID(c.Context(), req)
	if err != nil {
		return w.presenter.BuildError(c, err)
	}

	return w.presenter.BuildSuccess(c, nil, "Success", http.StatusOK)
}

// @Summary         Delete Todo List by ID
// @Description     Delete an existing Todo List by its ID
// @Tags			Todo List
// @Accept			json
// @Produce			json
// @Security 		Bearer
// @Param           id path int true "ID of the todo list"
// @Success			201 {object} entity.GeneralResponse "Success"
// @Failure			401 {object} entity.CustomErrorResponse "Unauthorized"
// @Failure			422 {object} entity.CustomErrorResponse "Invalid Request Body"
// @Failure			500 {object} entity.CustomErrorResponse "Internal server Error"
// @Router			/api/v1/todo-list-category/{id} [delete]
func (w *TodoListCategoryHandler) Delete(c *fiber.Ctx) error {
	id, err := w.parser.ParserIntIDFromPathParams(c)
	if err != nil {
		return w.presenter.BuildError(c, err)
	}

	err = w.todoListCategoryCrudUsecase.DeleteByID(c.Context(), id)
	if err != nil {
		return w.presenter.BuildError(c, err)
	}

	return w.presenter.BuildSuccess(c, nil, "Success", http.StatusOK)
}
