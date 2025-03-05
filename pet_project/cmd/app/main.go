package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"golang/pet_project/internal/database"
	"golang/pet_project/internal/handlers"
	"golang/pet_project/internal/tasksService"
	"golang/pet_project/internal/web/tasks"
)

func main() {
	database.InitDB()

	repo := tasksService.NewTaskRepository(database.DB)
	service := tasksService.NewService(repo)

	handler := handlers.NewHandler(service)
	
	// Инициализируем echo
	e := echo.New()
	
	// используем Logger и Recover
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Прикол для работы в echo. Передаем и регистрируем хендлер в echo
	strictHandler := tasks.NewStrictHandler(handler, nil)
	tasks.RegisterHandlers(e, strictHandler)

	if err := e.Start(":8080"); err != nil {
		log.Fatalf("failed to start with err: %v", err)
	}
}