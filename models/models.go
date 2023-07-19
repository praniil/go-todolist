package models

type TodoList struct{
	ID	int64 `json:"id"`
	Title	string `json:"title"`
	Status	string `json:"status"`
	CreatedAt	string `json:"createdat"`
}