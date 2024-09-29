package handlers

import (
	"encoding/json"
	"go-todo-app/constant"
	"go-todo-app/database"
	"go-todo-app/models"
	"go-todo-app/models/request"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

// GetTodos godoc
// @Summary Get todos
// @Description Get todos from related database
// @Tags todos
// @Accept  json
// @Produce  json
// @Success 200 {array} models.Todo "OK"
// @Failure 404 {string} string "Todo not found"
// @Failure 500 {string} string "Internal Server Error"
// @Router /GetTodos [get]
func GetTodos(w http.ResponseWriter, r *http.Request) {
	var todos []models.Todo

	result := database.DB.Find(&todos)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			w.WriteHeader(http.StatusNotFound)
			w.Header().Set("Content-Type", constant.ContentTypeJSON)
			json.NewEncoder(w).Encode(map[string]string{"error": "Todo not found"})
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			w.Header().Set("Content-Type", constant.ContentTypeJSON)
			json.NewEncoder(w).Encode(map[string]string{"error": "Error retrieving todo"})
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", constant.ContentTypeJSON)
	encoder := json.NewEncoder(w)
	encoder.SetIndent("", " ")
	encoder.Encode(todos)
}

// GetTodo godoc
// @Summary Get todo by ID
// @Description Get todo by ID from related database
// @Tags todos
// @Param id path int true "Todo ID"
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Todo "OK"
// @Failure 404 {string} string "Todo not found"
// @Failure 500 {string} string "Internal Server Error"
// @Router /GetTodo/{id} [get]
func GetTodo(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", constant.ContentTypeJSON)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid ID"})
		return
	}

	var todo models.Todo

	result := database.DB.First(&todo, id)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			w.WriteHeader(http.StatusNotFound)
			w.Header().Set("Content-Type", constant.ContentTypeJSON)
			json.NewEncoder(w).Encode(map[string]string{"error": "Todo not found"})
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			w.Header().Set("Content-Type", constant.ContentTypeJSON)
			json.NewEncoder(w).Encode(map[string]string{"error": "Error retrieving todo"})
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", constant.ContentTypeJSON)
	json.NewEncoder(w).Encode(todo)
}

// CreateNewTodo godoc
// @Summary Create a new todo
// @Description Create a new todo item in the database
// @Tags todos
// @Accept  json
// @Produce  json
// @Param todo body request.CreateTodoRequestModel true "Create New Todo Request"
// @Success 201 {object} models.Todo "Created"
// @Failure 400 {string} string "Bad request"
// @Failure 404 {string} string "Todo not found"
// @Failure 500 {string} string "Internal Server Error"
// @Router /CreateNewTodo [post]
func CreateNewTodo(w http.ResponseWriter, r *http.Request) {
	var todo models.Todo

	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if result := database.DB.Create(&todo); result.Error != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", constant.ContentTypeJSON)
	encoder := json.NewEncoder(w)
	encoder.SetIndent("", " ")
	encoder.Encode(todo)
}

// UpdateTodo godoc
// @Summary Update a todo by ID
// @Description Updates a todo by its ID from the database
// @Tags todos
// @Accept  json
// @Produce  json
// @Param todo body request.UpdateTodoRequestModel true "Update Todo Request"
// @Success 200 {object} models.Todo "OK"
// @Failure 400 {string} string "Bad request"
// @Failure 404 {string} string "Todo not found"
// @Failure 500 {string} string "Internal Server Error"
// @Router /UpdateTodo/{id} [put]
func UpdateTodo(w http.ResponseWriter, r *http.Request) {
	var updateRequest request.UpdateTodoRequestModel
	if err := json.NewDecoder(r.Body).Decode(&updateRequest); err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	var todo models.Todo
	result := database.DB.First(&todo, updateRequest.ID)
	if result.Error != nil {
		http.Error(w, "Todo not found", http.StatusNotFound)
		return
	}

	todo.Title = updateRequest.Title
	todo.Content = updateRequest.Content
	todo.Status = updateRequest.Status

	result = database.DB.Save(&todo)
	if result.Error != nil {
		http.Error(w, "Failed to update todo", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", constant.ContentTypeJSON)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(todo)
}

// DeleteTodo godoc
// @Summary Delete a todo by ID
// @Description Deletes a todo by its ID from the database
// @Tags todos
// @Param id path int true "Todo ID"
// @Success 204 "No Content"
// @Failure 400 {string} string "Bad request"
// @Failure 404 {string} string "Todo not found"
// @Failure 500 {string} string "Internal Server Error"
// @Router /DeleteTodo/{id} [delete]
func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var todo models.Todo
	result := database.DB.Delete(&todo, id)

	if result.RowsAffected == 0 {
		http.Error(w, "Todo not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// GetAllRecordsIncludeSoftDeleteds godoc
// @Summary Get all records include soft deleteds
// @Description Gets all records with soft deleted records.
// @Tags todos
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Todo "OK"
// @Failure 400 {string} string "Bad request"
// @Failure 404 {string} string "Todo not found"
// @Failure 500 {string} string "Internal Server Error"
// @Router /GetAllTodosIncludeSoftDeleteds [get]
func GetAllRecordsIncludeSoftDeleteds(w http.ResponseWriter, r *http.Request) {
	var todos []models.Todo
	database.DB.Unscoped().Find(&todos)

	w.Header().Set("Content-Type", constant.ContentTypeJSON)
	encoder := json.NewEncoder(w)
	encoder.SetIndent("", " ")
	encoder.Encode(todos)
}
