package router

import (
	"go-todoapp/handler"

	"github.com/gorilla/mux"
)

//Router is exported and used in main.go

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api/newtodos", handler.CreateTodolist).Methods("POST", "OPTIONS")
	return router
}
