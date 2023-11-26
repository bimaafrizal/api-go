package controller

import (
	"belajar-api-go/helper"
	"database/sql"
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"net/http"
)

type ToDoControllerImpl struct {
	DB       *sql.DB
	Validate *validator.Validate
}

func NewToDoControllerImpl(db *sql.DB, validate *validator.Validate) *ToDoControllerImpl {
	return &ToDoControllerImpl{DB: db, Validate: validate}
}

func (controller *ToDoControllerImpl) Create(ctx echo.Context) error {
	var body helper.CreateRequest
	err := controller.Validate.Struct(body)
	if err != nil {
		return helper.BadRequestResponse(err.Error())
	}

	err = json.NewDecoder(ctx.Request().Body).Decode(&body)
	if err != nil {
		return helper.InternalServerErrorResponse(err.Error())
	}

	_, err = controller.DB.Exec("INSERT INTO todos (title, description, done) VALUES (?, ?, 0)", body.Title, body.Description)
	if err != nil {
		return helper.InternalServerErrorResponse(err.Error())
	}

	return ctx.JSON(http.StatusOK, body)
}

func (controller *ToDoControllerImpl) GetAll(ctx echo.Context) error {
	rows, err := controller.DB.Query("SELECT id, title, description, done FROM todos")
	if err != nil {
		return helper.InternalServerErrorResponse(err.Error())
	}
	defer rows.Close()

	var todos []helper.ToDoResponse
	for rows.Next() {
		var todo helper.ToDoResponse
		var doneInt int

		err := rows.Scan(&todo.Id, &todo.Title, &todo.Description, &doneInt)
		if err != nil {
			return helper.InternalServerErrorResponse(err.Error())
		}

		if doneInt == 1 {
			todo.Done = true
		} else {
			todo.Done = false
		}

		todos = append(todos, todo)
	}

	return ctx.JSON(http.StatusOK, todos)
}

func (controller *ToDoControllerImpl) GetById(ctx echo.Context) error {
	id := ctx.Param("id")
	var todo helper.ToDoResponse

	err := controller.DB.QueryRow("SELECT id, title, description, done FROM todos WHERE id = ?", id).Scan(&todo.Id, &todo.Title, &todo.Description, &todo.Done)
	if err == sql.ErrNoRows {
		return helper.NotFoundErrorResponse()
	}
	if err != nil {
		return helper.InternalServerErrorResponse(err.Error())
	}

	return ctx.JSON(200, todo)
}

func (controller *ToDoControllerImpl) Update(ctx echo.Context) error {
	id := ctx.Param("id")
	var body helper.UpdateRequest
	err := controller.Validate.Struct(body)
	if err != nil {
		return helper.BadRequestResponse(err.Error())
	}

	err = json.NewDecoder(ctx.Request().Body).Decode(&body)
	if err != nil {
		return helper.InternalServerErrorResponse(err.Error())
	}

	err = controller.DB.QueryRow("SELECT id FROM todos WHERE id = ?", id).Scan(&id)
	if err == sql.ErrNoRows {
		return helper.NotFoundErrorResponse()
	}

	_, err = controller.DB.Exec("UPDATE todos SET title = ?, description = ? WHERE id = ?", body.Title, body.Description, id)
	if err != nil {
		return helper.InternalServerErrorResponse(err.Error())
	}

	return ctx.JSON(200, map[string]string{
		"message": "success update todo",
	})
}

func (controller *ToDoControllerImpl) Delete(ctx echo.Context) error {
	id := ctx.Param("id")

	err := controller.DB.QueryRow("SELECT id FROM todos WHERE id = ?", id).Scan(&id)
	if err == sql.ErrNoRows {
		return helper.NotFoundErrorResponse()
	}

	_, err = controller.DB.Exec("DELETE FROM todos WHERE id = ?", id)
	if err != nil {
		return helper.InternalServerErrorResponse(err.Error())
	}

	return ctx.JSON(200, map[string]string{
		"message": "success delete todo",
	})
}

func (controller *ToDoControllerImpl) Check(ctx echo.Context) error {
	id := ctx.Param("id")
	var body helper.CheckRequest
	err := controller.Validate.Struct(body)
	if err != nil {
		return helper.BadRequestResponse(err.Error())
	}
	err = json.NewDecoder(ctx.Request().Body).Decode(&body)
	if err != nil {
		return helper.InternalServerErrorResponse(err.Error())
	}

	var doneInt int
	if body.Done {
		doneInt = 1
	}

	err = controller.DB.QueryRow("SELECT id FROM todos WHERE id = ?", id).Scan(&id)
	if err == sql.ErrNoRows {
		return helper.NotFoundErrorResponse()
	}

	_, err = controller.DB.Exec("UPDATE todos SET done = ? WHERE id = ?", doneInt, id)
	if err != nil {
		return helper.InternalServerErrorResponse(err.Error())
	}
	return ctx.JSON(200, map[string]string{
		"message": "success check todo",
	})
}
