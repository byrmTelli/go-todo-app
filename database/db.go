package database

import (
	"go-todo-app/models"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDatabase() {
	var err error
	DB, err = gorm.Open(sqlite.Open("todos.db"), &gorm.Config{})

	if err != nil {
		log.Fatal("Error occurred while connecting to DB.", err)
	}

	DB.AutoMigrate(&models.Todo{}, &models.User{})
	log.Println("Database connection established and database migrated.")
}
