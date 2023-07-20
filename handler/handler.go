package handler

import (
	"encoding/json"
	"fmt"
	"go-todoapp/models"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Response struct {
	ID      int64  `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
}

var Database *gorm.DB

func Database_connection() *gorm.DB {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error in loading .env file")
	}

	host := os.Getenv("DB_HOST")
	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	database_name := os.Getenv("DB_NAME")
	port, err := strconv.Atoi(os.Getenv("DB_PORT"))

	if err != nil {
		log.Fatal()
	}

	psql_info := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable password=%s", host, port, username, database_name, password)

	Database, err := gorm.Open(postgres.Open(psql_info), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		log.Fatal()
	}

	return Database

}

func CreateTodolist(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	var todolist models.TodoList

	err := json.NewDecoder(r.Body).Decode(&todolist)
	if err != nil {
		log.Fatalf("unable to decode the request body, %v", err)
	}
	insertID := inserttodolist(todolist)
	res := Response{
		ID:      insertID,
		Message: "todos successfully created",
	}

	json.NewEncoder(w).Encode(res)

}

func inserttodolist(todolist models.TodoList) int64{
	db := Database_connection()
	db.AutoMigrate(&models.TodoList{})
	result := db.Create(&todolist)
	if result.Error != nil{
		panic(fmt.Sprintf("unable to create todolist"))
	}
	fmt.Printf("inserted a single todos, %v",todolist.ID)
	return todolist.ID
}
