package services

import (
	"github.com/ankitanwar/GoAPIUtils/errors"
	"github.com/ankitanwar/e-Commerce/User/domain/users"
)

var (
	//UserServices : All the services available for user
	UserServices userServicesInterface = &userServices{}
)

type userServices struct{}

type userServicesInterface interface {
	CreateUser(users.User) (*users.User, *errors.RestError)
	GetUser(int) (*users.User, *errors.RestError)
	UpdateUser(bool, users.User) (*users.User, *errors.RestError)
	DeleteUser(int) *errors.RestError
	FindByStatus(string) (users.Users, *errors.RestError)
	LoginUser(request users.LoginRequest) (*users.User, *errors.RestError)
}

//CreateUser : To save the user in the database
func (u *userServices) CreateUser(newUser users.User) (*users.User, *errors.RestError) {
	if err := newUser.Validate(); err != nil {
		return nil, err
	}
	if err := newUser.Save(); err != nil {
		return nil, err
	}
	return &newUser, nil
}

//GetUser : To get the detail of the user with given id
func (u *userServices) GetUser(userid int) (*users.User, *errors.RestError) {
	if userid < 0 {
		return nil, errors.NewBadRequest("Enter the valied user id")
	}
	result := &users.User{
		ID: userid,
	}
	if err := result.Get(); err != nil {
		return nil, err
	}
	return result, nil
}

//UpdateUser : To update the values of the existing user
func (u *userServices) UpdateUser(partial bool, user users.User) (*users.User, *errors.RestError) {
	current := users.User{
		ID: user.ID,
	}
	if err := current.Get(); err != nil {
		return nil, err
	}
	if err := user.Validate(); err != nil {
		return nil, err
	}
	if partial {
		if user.FirstName != "" {
			current.FirstName = user.FirstName
		}
		if user.LastName != "" {
			current.LastName = user.LastName
		}
		if user.Email != "" {
			current.Email = user.Email
		}
	} else {
		current.FirstName = user.FirstName
		current.LastName = user.LastName
		current.Email = user.Email
	}
	if err := current.Update(); err != nil {
		return nil, err
	}
	return &current, nil
}

//DeleteUser : To delete the given user
func (u *userServices) DeleteUser(userID int) *errors.RestError {
	user := &users.User{
		ID: userID,
	}
	return user.Delete()
}

//FindByStatus : To find the user by status
func (u *userServices) FindByStatus(status string) (users.Users, *errors.RestError) {
	users := users.User{}

	foundUsers, err := users.FindByStatus(status)
	if err != nil {
		return nil, err
	}
	return foundUsers, nil

}

func (u *userServices) LoginUser(request users.LoginRequest) (*users.User, *errors.RestError) {
	user := &users.User{}
	user.Email = request.Email
	user.Password = request.Password
	if err := user.GetUserByEmailAndPassword(); err != nil {
		return nil, err
	}
	return user, nil
}
