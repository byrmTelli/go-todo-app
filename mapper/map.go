package mapper

import (
	"go-todo-app/models"
	"go-todo-app/models/response"
)

func UserDTOMapper(user models.User) response.UserDTO {
	return response.UserDTO{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
	}
}

func UserToLoginDTOMapper(user models.User, token string) response.UserLoginDTO {
	return response.UserLoginDTO{
		ID:          user.ID,
		Username:    user.Username,
		Email:       user.Email,
		AccessToken: token,
	}
}
