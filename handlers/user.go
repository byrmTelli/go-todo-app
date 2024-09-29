package handlers

import (
	"encoding/json"
	"go-todo-app/constant"
	"go-todo-app/database"
	"go-todo-app/helpers"
	"go-todo-app/mapper"
	"go-todo-app/models"
	"go-todo-app/models/response"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

// GetUsers godoc
// @Summary Get all users
// @Description Get all users
// @Tags users
// @Accept  json
// @Produce  json
// @Success 200 {object} models.User "OK"
// @Failure 400 {string} string "Bad request"
// @Failure 404 {string} string "Todo not found"
// @Failure 500 {string} string "Internal Server Error"
// @Router /GetUsers [get]
func GetUsers(w http.ResponseWriter, r *http.Request) {
	var users []models.User
	database.DB.Find(&users)

	var userDTOs []response.UserDTO

	for _, user := range users {
		userDTO := mapper.UserDTOMapper(user)
		userDTOs = append(userDTOs, userDTO)
	}

	w.Header().Set("Content-Type", constant.ContentTypeJSON)
	encoder := json.NewEncoder(w)
	encoder.Encode(userDTOs)
}

// Register godoc
// @Summary Create a new user
// @Description Create a new user with the given details
// @Tags users
// @Accept  json
// @Produce  json
// @Param user body request.RegisterRequestModel true "Create User Request"
// @Success 200 {object} string "OK"
// @Failure 400 {string} string "Bad request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /Register [post]
func Register(w http.ResponseWriter, r *http.Request) {
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)

	hashedPassword, err := helpers.HashPassword(user.Password)
	if err != nil {
		http.Error(w, "Error while hashing password.", http.StatusInternalServerError)
	}

	user.Password = hashedPassword

	database.DB.Create(&user)

	w.Header().Set("Content-Type", constant.ContentTypeJSON)
	encoder := json.NewEncoder(w)
	encoder.Encode(map[string]string{"message": "User created succesfully"})
}

// GetUser godoc
// @Summary Get a user by ID
// @Description Get user details by their ID
// @Tags users
// @Accept  json
// @Produce  json
// @Param id path int true "User ID"
// @Success 200 {object} models.User "OK"
// @Failure 400 {string} string "Bad request"
// @Failure 404 {string} string "Todo not found"
// @Failure 500 {string} string "Internal Server Error"
// @Router /GetUser/{id} [get]
func GetUser(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", constant.ContentTypeJSON)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid value."})
		return
	}

	var user models.User

	result := database.DB.First(&user, id)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			w.WriteHeader(http.StatusNotFound)
			w.Header().Set("Content-Type", constant.ContentTypeJSON)
			json.NewEncoder(w).Encode(map[string]string{"error": "User Not Found."})

		} else {
			w.WriteHeader(http.StatusInternalServerError)
			w.Header().Set("Content-Type", constant.ContentTypeJSON)
			json.NewEncoder(w).Encode(map[string]string{"error": "Error while retrieving user data."})
		}
		return
	}

	userDTO := mapper.UserDTOMapper(user)

	w.Header().Set("Content-Type", constant.ContentTypeJSON)
	json.NewEncoder(w).Encode(userDTO)
}
