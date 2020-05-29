package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"log"
	"os"
)

var db *gorm.DB

func init() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
	DBUser := os.Getenv("DB_USER_NAME")
	DBPass := os.Getenv("DB_PASSWORD")
	DBHost := os.Getenv("DB_HOST")
	DBPort := os.Getenv("DB_PORT")
	DBName := os.Getenv("DB_DBNAME")
	dataSource := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=require",
		DBHost, DBPort, DBUser, DBName, DBPass)
	db, err = gorm.Open("postgres", dataSource)
	if err != nil {
		fmt.Println(err)
	}
	db.AutoMigrate(Record{}, Author{}, AsyncManage{})
	db.LogMode(true)
}

type ApiResponse struct {
	Status  string
	Message string
}

func TruncateAuthor() {
	db.Exec("DELETE FROM authors")
}

func TruncateRecord() {
	db.Exec("DELETE FROM records")
}

func OutputLogIfError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
