package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type requestBody struct {
	Task string `json:"message"`
}

var task string

func GetHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		fmt.Fprintln(w, "hello,", task)
	} else {
		fmt.Fprintln(w, "Only GET")
	}
}

func PostHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var rB requestBody
		json.NewDecoder(r.Body).Decode(&rB)
		task = rB.Task
	} else {
		fmt.Fprintln(w, "Only POST")
	}
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/get", GetHandler)
	router.HandleFunc("/post", PostHandler)
	http.ListenAndServe(":8080", router)
}
