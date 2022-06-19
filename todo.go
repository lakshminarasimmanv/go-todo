/* todo app using Go. */

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB
var err error

func main() {
	db, err = gorm.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	db.AutoMigrate(&todoModel{})

	router := mux.NewRouter()
	router.HandleFunc("/todos", getTodos).Methods("GET")
	router.HandleFunc("/todos", createTodo).Methods("POST")
	router.HandleFunc("/todos/{id}", getTodo).Methods("GET")
	router.HandleFunc("/todos/{id}", updateTodo).Methods("PUT")
	router.HandleFunc("/todos/{id}", deleteTodo).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
}

type todoModel struct {
	gorm.Model
	Title     string    `json:"title"`
	Completed bool      `json:"completed"`
	Due       time.Time `json:"due"`
}

type todos []todoModel

func getTodos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var todos []todoModel
	db.Find(&todos)
	json.NewEncoder(w).Encode(todos)
}

func createTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var todo todoModel
	_ = json.NewDecoder(r.Body).Decode(&todo)
	db.Create(&todo)
	json.NewEncoder(w).Encode(todo)
}

func getTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var todo todoModel
	db.First(&todo, params["id"])
	json.NewEncoder(w).Encode(todo)
}

func updateTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var todo todoModel
	db.First(&todo, params["id"])
	_ = json.NewDecoder(r.Body).Decode(&todo)
	db.Save(&todo)
	json.NewEncoder(w).Encode(todo)
}

func deleteTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var todo todoModel
	db.First(&todo, params["id"])
	db.Delete(&todo)
	json.NewEncoder(w).Encode(todo)
}
