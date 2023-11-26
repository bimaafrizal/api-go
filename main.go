package main

import (
	"belajar-api-go/controller"
	"belajar-api-go/database"
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
)

func main() {
	db := database.InitDB()
	validate := validator.New()
	defer db.Close()
	err := db.Ping()
	userController := controller.NewToDoControllerImpl(db, validate)

	if err != nil {
		panic(err)
	}

	e := echo.New()
	e.POST("todos", userController.Create)
	e.GET("todos", userController.GetAll)
	e.GET("todos/:id", userController.GetById)
	e.PATCH("todos/:id", userController.Update)
	e.DELETE("todos/:id", userController.Delete)
	e.PATCH("todos/:id/check", userController.Check)
	e.Start(":8080")

	if err != nil {
		panic(err)
	}
}
