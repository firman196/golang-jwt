package users

import (
	"golang-jwt/model/entity"
)

type UsersResponse struct {
	Id        int16  `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
}

func UsersResponses(user entity.Users) UsersResponse {
	return UsersResponse{
		Id:        user.Id,
		Firstname: user.Firstname,
		Lastname:  user.Lastname,
		Email:     user.Email,
	}
}
