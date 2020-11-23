package users

//LoginRequest : To get the login request from the user
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
