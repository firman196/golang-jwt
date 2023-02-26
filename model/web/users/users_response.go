package users

import (
	"golang-jwt/model/entity"
)

type UsersResponse struct {
	Id        int16  `json:"id"`
	Firstname string `json:"first_name"`
	Lastname  string `json:"last_name"`
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
