package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"user-service/internal/database"
	"user-service/internal/handler"
)

func main() {
	database.InitDB()

	r := mux.NewRouter()

	r.HandleFunc("/users", handler.CreateUser).Methods("POST")
	r.HandleFunc("/users/{id}", handler.GetUser).Methods("GET")
	r.HandleFunc("/users/{id}", handler.UpdateUser).Methods("PUT")
	r.HandleFunc("/users/{id}", handler.DeleteUser).Methods("DELETE")
	r.HandleFunc("/users", handler.GetAllUsers).Methods("GET")

	log.Println("Server starting on :5000")
	log.Fatal(http.ListenAndServe(":5000", r))
}