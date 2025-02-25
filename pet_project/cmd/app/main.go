package main

import (
	"github.com/gorilla/mux"
	"golang/pet_project/internal/database"
	"golang/pet_project/internal/handlers"
	"golang/pet_project/internal/taskService"
	"net/http"
)

func main() {
	database.InitDB()
	database.DB.AutoMigrate(&taskService.Task{})

	repo := taskService.NewTaskRepository(database.DB)
	service := taskService.NewService(repo)

	handler := handlers.NewHandler(service)

	router := mux.NewRouter()
	router.HandleFunc("/api/get", handler.GetTasksHandler).Methods("GET")
	router.HandleFunc("/api/post", handler.PostTaskHandler).Methods("POST")
	router.HandleFunc("/api/patch/{id}", handler.UpdateTaskByID).Methods("PATCH")
	router.HandleFunc("/api/delete/{id}", handler.DeleteTaskByID).Methods("DELETE")
	http.ListenAndServe(":8080", router)
}
