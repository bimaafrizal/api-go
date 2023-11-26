package controller

import "github.com/labstack/echo/v4"

type ToDoController interface {
	Create(ctx echo.Context) error
	GetAll(ctx echo.Context) error
	GetById(ctx echo.Context) error
	Update(ctx echo.Context) error
	Delete(ctx echo.Context) error
	Check(ctx echo.Context) error
}
