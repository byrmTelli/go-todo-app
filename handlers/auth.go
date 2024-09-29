package handlers

import (
	"encoding/json"
	"go-todo-app/constant"
	"go-todo-app/database"
	"go-todo-app/helpers"
	"go-todo-app/mapper"
	"go-todo-app/models"
	"go-todo-app/models/request"
	"net/http"
)

// Login godoc
// @Summary User login
// @Description Login a user with username and password
// @Tags auth
// @Accept  json
// @Produce  json
// @Param LoginRequestModel body request.LoginRequestModel true "Login Request"
// @Success 200 {object} response.UserLoginDTO
// @Failure 400 {string} string "Invalid input"
// @Failure 401 {string} string "Unauthorized"
// @Failure 500 {string} string "Internal Server Error"
// @Router /login [post]
func Login(w http.ResponseWriter, r *http.Request) {
	var loginRequestModel request.LoginRequestModel

	// Decoding Body
	err := json.NewDecoder(r.Body).Decode(&loginRequestModel)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	// Check Username & Password Status
	if loginRequestModel.Username == "" || loginRequestModel.Password == "" {
		http.Error(w, "Username and Password fields are required.", http.StatusBadRequest)
		return
	}

	// Check is user exist
	var user models.User
	result := database.DB.Where("username = ?", loginRequestModel.Username).First(&user)
	if result.Error != nil {
		http.Error(w, "Username or password wrong.", http.StatusUnauthorized)
		return
	}

	// Compare request's password with user's hash.
	if !helpers.CheckPasswordHash(loginRequestModel.Password, user.Password) {
		http.Error(w, "username or password wrong.", http.StatusUnauthorized)
		return
	}

	token, err := GenerateJWT(user)
	if err != nil {
		http.Error(w, "Error occured while generating token", http.StatusInternalServerError)
		return
	}

	var userLoginDTO = mapper.UserToLoginDTOMapper(user, token)

	// Success Result
	w.Header().Set("Content-Type", constant.ContentTypeJSON)
	encoder := json.NewEncoder(w)
	encoder.SetIndent("", " ")
	encoder.Encode(userLoginDTO)
}
