package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func GetHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		var tasks []Task
		DB.Find(&tasks)
		w.Header().Set("Content-Type", "application/json") // Установка формата
		json.NewEncoder(w).Encode(tasks)                   // Декодирование
	} else {
		fmt.Fprintln(w, "Only GET")
	}
}

func PostHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var task Task
		json.NewDecoder(r.Body).Decode(&task)
		DB.Create(&task)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(task)
	} else {
		fmt.Fprintln(w, "Only POST")
	}
}

func PatchHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPatch {
		vars := mux.Vars(r)
		taskID := vars["id"]
		var task Task
		DB.First(&task, taskID) // поиск задачи по ключу
		var update map[string]interface{}
		json.NewDecoder(r.Body).Decode(&update)
		DB.Model(&task).Updates(update) // Обновление нескольких полей
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(task)
	} else {
		fmt.Fprintln(w, "Only PATCH")
	}
}

func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodDelete {
		vars := mux.Vars(r)
		taskID := vars["id"]
		var task Task
		DB.Delete(&task, taskID)
	} else {
		fmt.Fprintln(w, "Only DELETE")
	}

}
func main() {
	InitDB()
	DB.AutoMigrate(&Task{})

	router := mux.NewRouter()
	router.HandleFunc("/get", GetHandler).Methods("GET")
	router.HandleFunc("/post", PostHandler).Methods("POST")
	router.HandleFunc("/patch/{id}", PatchHandler).Methods("PATCH")
	router.HandleFunc("/delete/{id}", DeleteHandler).Methods("DELETE")
	http.ListenAndServe(":8080", router)
}
