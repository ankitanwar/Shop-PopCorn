package users

import (
	"strings"

	"github.com/ankitanwar/GoAPIUtils/errors"
)

//User : User and its values
type User struct {
	ID          int    `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
	Status      string `json:"status"`
	Password    string `json:"password"`
}

//Users : It will return the slices of users
type Users []User

//Validate : To validate the users
func (user *User) Validate() *errors.RestError {
	if user.FirstName == "" {
		err := errors.NewBadRequest("Please Enter the First Name")
		return err
	}
	if user.LastName == "" {
		err := errors.NewBadRequest("Please Enter the Last Name")
		return err
	}
	if user.Email == "" {
		err := errors.NewBadRequest("Please enter the valid mail address")
		return err
	}
	user.Password = strings.TrimSpace(user.Password)
	if user.Password == "" {
		return errors.NewBadRequest("Please Enter the valid password")
	}
	return nil
}
