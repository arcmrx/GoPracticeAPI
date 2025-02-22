package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func GetHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		var tasks []Task
		DB.Find(&tasks)
		w.Header().Set("Content-Type", "application/json") // Установка формата
		json.NewEncoder(w).Encode(tasks) // Декодирование 
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

func main() {
	InitDB()
	DB.AutoMigrate(&Task{})

	router := mux.NewRouter()
	router.HandleFunc("/get", GetHandler).Methods("GET")
	router.HandleFunc("/post", PostHandler).Methods("POST")
	http.ListenAndServe(":8080", router)
}
