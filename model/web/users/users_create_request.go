package users

type UsersCreateRequest struct {
	Firstname string `validate:"required, min=1,max=100" json:"first_name"`
	Lastname  string `validate:"required, min=1, max=100" json:"last_name"`
	Email     string `validate:"required, min=1, max=100" json:"email"`
	Password  string `validate:"required" json:"password"`
}
