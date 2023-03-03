package users

type UsersCreateRequest struct {
	Firstname string `validate:"required,min=1,max=100" json:"firstname"`
	Lastname  string `validate:"required,min=1,max=100" json:"lastname"`
	Email     string `validate:"required,min=1,max=100" json:"email"`
	Password  string `validate:"required" json:"password"`
}
