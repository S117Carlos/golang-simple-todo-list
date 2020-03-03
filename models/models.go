package models

type TodoItem struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	IsChecked   bool   `json:"isChecked"`
}

type TodosList []TodoItem
