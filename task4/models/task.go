package models

type Task struct {
	ID int `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	DueDate string `json:"due_date"`
	Done string `json:"done"`
}
