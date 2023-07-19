package handler

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"os"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
)

type Response struct{
	ID int64 `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
}

var Database *gorm.DB
func Database_connection() *gorm.DB{
	err:= godotenv.Load(".env")

	if err!=nil{
		log.Fatal("Error in loading .env file")
	}

	host := os.Getenv("DB_HOST")
	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	database_name := os.Getenv("DB_NAME")
	port, err := strconv.Atoi(os.Getenv("DB_PORT"))

	if err != nil{
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

func CreateTodolist(w http.ResponseWriter, r *http.Request){

}