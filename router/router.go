package router

import (
	"go-todoapp/handler"

	"github.com/gorilla/mux"
)

//Router is exported and used in main.go

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api/newtodos", handler.CreateTodolist).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/gettodo/{id}", handler.GetTodolisit).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/updatetodo", handler.UpdateTodolist).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/getalltodos", handler.GetAllTodolist).Methods("GET", "OPTIONS")
	return router
}
