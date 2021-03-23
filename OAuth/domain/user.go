package domain

//User : attributes of user
type User struct {
	UserID    string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

//LoginRequest : To request for login from the given email and password
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
