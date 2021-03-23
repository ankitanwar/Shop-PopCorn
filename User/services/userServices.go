package services

import (
	"github.com/ankitanwar/GoAPIUtils/errors"
	userDB "github.com/ankitanwar/Shop-PopCorn/User/databasource/sql"
	"github.com/ankitanwar/Shop-PopCorn/User/domain/users"
)

var (
	//UserServices : All the services available for user
	UserServices userServicesInterface = &userServices{}
)

type userServices struct{}

type userServicesInterface interface {
	CreateUser(users.User) (*users.User, *errors.RestError)
	GetUser(string) (*users.User, *errors.RestError)
	UpdateUser(users.User) (*users.User, *errors.RestError)
	DeleteUser(string) *errors.RestError
	VerifyUser(request users.LoginRequest) (*users.User, *errors.RestError)
}

//CreateUser : To save the user in the database
func (u *userServices) CreateUser(newUser users.User) (*users.User, *errors.RestError) {
	if err := newUser.Validate(); err != nil {
		return nil, err
	}
	if err := userDB.Save(&newUser); err != nil {
		return nil, err
	}
	return &newUser, nil
}

//GetUser : To get the detail of the user with given id
func (u *userServices) GetUser(userid string) (*users.User, *errors.RestError) {
	if userid == "" {
		return nil, errors.NewBadRequest("Enter the valied user id")
	}
	find, err := userDB.Get(userid)
	if err != nil {
		return nil, err
	}
	return find, nil
}

//UpdateUser : To update the values of the existing user
func (u *userServices) UpdateUser(user users.User) (*users.User, *errors.RestError) {
	savedDetails, err := userDB.Get(user.ID)
	if err != nil {
		return nil, err
	}
	if user.FirstName != "" {
		savedDetails.FirstName = user.FirstName
	}
	if user.LastName != "" {
		savedDetails.LastName = user.LastName
	}
	if user.Email != "" {
		savedDetails.Email = user.Email
	}
	if user.PhoneNo != "" {
		if len(user.PhoneNo) < 10 || len(user.PhoneNo) > 10 {
			return nil, errors.NewBadRequest("Enter The valid Phone Number")
		}
	}
	if err := userDB.Update(savedDetails); err != nil {
		return nil, err
	}
	return savedDetails, nil
}

//DeleteUser : To delete the given user
func (u *userServices) DeleteUser(userID string) *errors.RestError {
	err := userDB.Delete(userID)
	if err != nil {
		return err
	}
	return nil
}

func (u *userServices) VerifyUser(request users.LoginRequest) (*users.User, *errors.RestError) {
	user := &users.User{}
	user.Email = request.Email
	user.Password = request.Password
	if err := userDB.GetUserByEmailAndPassword(user); err != nil {
		return nil, err
	}
	return user, nil
}
