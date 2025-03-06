package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"golang/pet_project/internal/database"
	"golang/pet_project/internal/handlers"
	"golang/pet_project/internal/tasksService"
	"golang/pet_project/internal/userService"
	"golang/pet_project/internal/web/tasks"
	"golang/pet_project/internal/web/users"
)

func main() {
	database.InitDB()

	tasksRepo := tasksService.NewTaskRepository(database.DB)
	taskService := tasksService.NewTaskService(tasksRepo)
	taskHandler := handlers.NewTaskHandler(taskService)

	usersRepo := userservice.NewUserRepository(database.DB)
	userService := userservice.NewUserService(usersRepo)
	userHandler := handlers.NewUserHandler(userService)
	
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	strictTaskHandler := tasks.NewStrictHandler(taskHandler, nil)
	tasks.RegisterHandlers(e, strictTaskHandler)

	strictUserHandler := users.NewStrictHandler(userHandler, nil)
	users.RegisterHandlers(e, strictUserHandler)

	if err := e.Start(":8080"); err != nil {
		log.Fatalf("failed to start with err: %v", err)
	}
}