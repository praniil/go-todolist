package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"go-todoapp/models"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
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

func GetTodolisit(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Fatalf("unable to convert string to int")
	}
	todolist, err := gettodolist(int64(id))
	if err != nil {
		log.Fatalf("unable to get todolist")
	}
	json.NewEncoder(w).Encode(todolist)
}

func GetAllTodolist(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	todolist, err := getalltodolist()
	if err != nil{
		log.Fatalf("unable to get todolists, %v", err)
	}
	json.NewEncoder(w).Encode(todolist)

}

func UpdateTodolist(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	var todolist models.TodoList
	err := json.NewDecoder(r.Body).Decode(&todolist)
	if err!= nil{
		log.Fatalf("unable to decode request body, %v", err)
	}
	updatedRows := updatetodolist(todolist.ID, todolist)
	msg := fmt.Sprintf("the numbers of updated todolist: %v", updatedRows)
	res:= Response{
		ID : todolist.ID,
		Message : msg,
	}
	json.NewEncoder(w).Encode(res)
}

// func DeleteTodolist(w http.ResponseWriter, r *http.Request){
// 	w.Header().Set("Content-type", "application/x-www-form-urlencoded")
// 	w.Header().Set("Access-Control-Allow-Origin", "*")
// 	w.Header().Set("Access-Control-Allow-Methods", "PUT")
// 	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

// 	params := mux.Vars(r)
// 	id, err := strconv.Atoi(params["id"])
// 	if err!=nil{
// 		log.Fatal("unable to convert string to int")
// 	}
// 	deletedRows := deletetodolist(int64(id))
// 	msg := fmt.Sprintf("no of deleted rows:%%v", deletedRows)
// 	res:= Response{
// 		ID : int64(id),
// 		Message: msg,
// 	}
// 	json.NewEncoder(w).Encode(res)

// }

func inserttodolist(todolist models.TodoList) int64 {
	db := Database_connection()
	db.AutoMigrate(&models.TodoList{})
	result := db.Create(&todolist)
	if result.Error != nil {
		panic(fmt.Sprintf("unable to create todolist"))
	}
	fmt.Printf("inserted a single todos, %v", todolist.ID)
	return todolist.ID
}

func gettodolist(id int64) (models.TodoList, error) {
	db := Database_connection()
	var todolist models.TodoList
	result := db.First(&todolist, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		fmt.Println("no rows were returned")
		return todolist, nil
	} else if result.Error != nil {
		log.Fatalf("unable to query todos, %v", result.Error)
		return todolist, result.Error
	}
	return todolist, nil
}

func getalltodolist() ([]models.TodoList, error){
	db:= Database_connection()
	var todolists []models.TodoList
	result := db.Find(&todolists)
	if result.Error!=nil{
		log.Fatalf("unable to find todolist, %v", result.Error)
		return todolists, result.Error
	}
	return todolists, nil
}

func updatetodolist(id int64, todolist models.TodoList) int64{
	db:= Database_connection()
	result := db.Model(models.TodoList{}).Where("id =?", id).Updates(todolist)
	if result.Error != nil{
		log.Fatalf("unable to update todos, %v", result.Error)
	}
	rowsAffected := result.RowsAffected
	log.Printf("no of rows affected: %v", rowsAffected)
	return rowsAffected
}

// func deletetodolist(id int64) int64{
// 	db := Database_connection()
// 	result := db.Delete(&models.TodoList{}, id)
// 	if result.Error != nil{
// 		log.Fatalf("unable to delete todolist, %v", result.Error)
// 	}
// 	rowsDeleted := result.RowsAffected
// 	return rowsDeleted
// }