package users

//PublicUser : Public User Marshall Struct
type PublicUser struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

//PrivateUser : When user is interacting with itself or internally with the app
type ReturnUserDetails struct {
	UserID    string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

//ReturnAddress : To return the address of the user
type ReturnAddress struct {
	Address string `json:"address"`
	State   string `json:"state"`
	Country string `json:"country"`
	Phone   string `json:"phone"`
}

func (display *ReturnUserDetails) ShowDetails(user *User) {
	display.UserID = user.ID
	display.FirstName = user.FirstName
	display.LastName = user.LastName
	display.Email = user.Email
}
