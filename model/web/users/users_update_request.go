package users

type UsersUpdateRequest struct {
	Id        int16  `json:"id"`
	Firstname string `validate:"required,min=1,max=100" json:"firstname"`
	Lastname  string `validate:"required,min=1,max=100" json:"lastname"`
	Email     string `validate:"required,min=1,max=100,email,unique" json:"email"`
	Password  string `validate:"required" json:"password"`
}
