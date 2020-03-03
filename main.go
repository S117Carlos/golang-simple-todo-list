package main

import (
	"net/http"

	"github.com/S117Carlos/golang-simple-todo-list/handlers"
)

func main() {
	http.Handle("/", handlers.DecorateRequest(handlers.GetTasks))
	http.HandleFunc("/getTodos", handlers.GetTasks)
	http.HandleFunc("/markTodo", handlers.MarkTodo)
	http.HandleFunc("/addTodo", handlers.AddTodo)
	http.ListenAndServe(":8081", nil)
}
