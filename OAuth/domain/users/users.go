package users

//User : User strcut
type User struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last__name"`
	Email     string `json:"email"`
}

//LoginRequest : Struct to send login request from the user and validate the user
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
