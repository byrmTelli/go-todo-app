package routes

import (
	"go-todo-app/database"
	"go-todo-app/handlers"
	"go-todo-app/middlewares"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

func Run(port string) {
	database.InitDatabase()

	router := mux.NewRouter()

	// Swagger Docs
	router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	// Auth Handlers
	router.HandleFunc("/Login", handlers.Login).Methods("POST")

	// User Handlers
	router.Handle("/GetUsers", middlewares.JWTAuthorizationMiddleware(http.HandlerFunc(handlers.GetUsers))).Methods("GET")
	router.HandleFunc("/Register", handlers.Register).Methods("POST")
	router.Handle("/GetUser/{id}", middlewares.JWTAuthorizationMiddleware(http.HandlerFunc(handlers.GetUser))).Methods("GET")

	// Todo Handlers
	router.Handle("/GetListOfTodos", middlewares.JWTAuthorizationMiddleware(http.HandlerFunc(handlers.GetTodos))).Methods("GET")
	router.Handle("/GetTodo/{id}", middlewares.JWTAuthorizationMiddleware(http.HandlerFunc(handlers.GetTodo))).Methods("GET")
	router.Handle("/GetAllTodosIncludeSoftDeleteds", middlewares.JWTAuthorizationMiddleware(http.HandlerFunc(handlers.GetAllRecordsIncludeSoftDeleteds))).Methods("GET")
	router.Handle("/CreateNewTodo", middlewares.JWTAuthorizationMiddleware(http.HandlerFunc(handlers.CreateNewTodo))).Methods("POST")
	router.Handle("/UpdateTodo/{id}", middlewares.JWTAuthorizationMiddleware(http.HandlerFunc(handlers.UpdateTodo))).Methods("PUT")
	router.Handle("/DeleteTodo/{id}", middlewares.JWTAuthorizationMiddleware(http.HandlerFunc(handlers.DeleteTodo))).Methods("POST")

	log.Println("Server is runnning on port ", port)
	log.Fatal(http.ListenAndServe(port, router))
}
