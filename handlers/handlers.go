package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/S117Carlos/golang-simple-todo-list/models"
)

var todoList models.TodosList

var currentID int

func DecorateRequest(handler func(http.ResponseWriter, *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.Background()
		if r.Context() != nil {
			ctx = r.Context()
		}

		ctx, cancel := context.WithCancel(ctx)
		defer cancel()
		handler(w, r)
	})
}

func AddTodo(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	var item models.TodoItem
	err := json.NewDecoder(req.Body).Decode(&item)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	currentID := currentID + 1
	item.ID = currentID
	select {
	case <-time.After(2 * time.Second):
		todoList = append(todoList, item)
		fmt.Fprintf(w, "TODO ADDED: %v", item.Title)
	case <-ctx.Done():
		fmt.Fprintf(w, "TODO ADDED: %v", ctx.Err())
	}
}

func MarkTodo(w http.ResponseWriter, req *http.Request) {
	var item models.TodoItem
	err := json.NewDecoder(req.Body).Decode(&item)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for i := range todoList {
		if todoList[i].ID == item.ID {
			todoList[i].IsChecked = !todoList[i].IsChecked
		}
	}
}

func GetTasks(w http.ResponseWriter, req *http.Request) {
	// Declare a new Person struct.
	var item models.TodoItem

	// Try to decode the request body into the struct. If there is an error,
	// respond to the client with the error message and a 400 status code.
	err := json.NewDecoder(req.Body).Decode(&item)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for _, t := range todoList {
		if t.IsChecked {
			fmt.Fprintf(w, "[X]: %v", t.Title)
		} else {
			fmt.Fprintf(w, "[]: %v", t.Title)
		}
	}

}
