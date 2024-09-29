package main

import (
	_ "go-todo-app/docs"
	"go-todo-app/routes"
)

// @title Bayram Telli ToDo App
// @version 1.0.0
// @description This is a simple todo app api that created by Bayram TELLI to practice golang.
// @termsOfService http://swagger.io/terms/

// @contact.name Bayram TELLİ
// @contact.url http://www.samplemail.com
// @contact.email sample@mail.com

// @host localhost:8080
// @BasePath /
func main() {

	routes.Run(":8080")
}
