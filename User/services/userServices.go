package services

import (
	"github.com/ankitanwar/GoAPIUtils/errors"
	addressDB "github.com/ankitanwar/e-Commerce/User/databasource/mongoDB"
	userDB "github.com/ankitanwar/e-Commerce/User/databasource/postgres"
	"github.com/ankitanwar/e-Commerce/User/domain/users"
)

var (
	//UserServices : All the services available for user
	UserServices userServicesInterface = &userServices{}
)

type userServices struct{}

type userServicesInterface interface {
	CreateUser(users.User) (*users.User, *errors.RestError)
	GetUser(string) (*users.User, *errors.RestError)
	UpdateUser(bool, users.User) (*users.User, *errors.RestError)
	DeleteUser(string) *errors.RestError
	FindByStatus(string) (users.Users, *errors.RestError)
	LoginUser(request users.LoginRequest) (*users.User, *errors.RestError)
	GetAddress(string) (*users.Address, *errors.RestError)
	AddAddress(string, *users.UserAddress) *errors.RestError
	RemoveAddress(string)
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
func (u *userServices) UpdateUser(partial bool, user users.User) (*users.User, *errors.RestError) {
	current := &users.User{
		ID: user.ID,
	}
	find, err := userDB.Get(user.ID)
	if err != nil {
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
		current.FirstName = find.FirstName
		current.LastName = find.LastName
		current.Email = find.Email
	}
	if err := userDB.Update(current); err != nil {
		return nil, err
	}
	return current, nil
}

//DeleteUser : To delete the given user
func (u *userServices) DeleteUser(userID string) *errors.RestError {
	err := userDB.Delete(userID)
	if err != nil {
		return err
	}
	return nil
}

//FindByStatus : To find the user by status
func (u *userServices) FindByStatus(status string) (users.Users, *errors.RestError) {
	foundUsers, err := userDB.FindByStatus(status)
	if err != nil {
		return nil, err
	}
	return foundUsers, nil

}

func (u *userServices) LoginUser(request users.LoginRequest) (*users.User, *errors.RestError) {
	user := &users.User{}
	user.Email = request.Email
	user.Password = request.Password
	if err := userDB.GetUserByEmailAndPassword(user); err != nil {
		return nil, err
	}
	return user, nil
}

func (u *userServices) GetAddress(userID string) (*users.Address, *errors.RestError) {
	find, err := addressDB.GetUserAddress(userID)
	if err != nil {
		return nil, errors.NewInternalServerError("Error While fetching the address")
	}
	return find, nil
}

func (u *userServices) AddAddress(userID string, address *users.UserAddress) *errors.RestError {
	address.ValidateAddress()
	err := addressDB.AddAddress(userID, address)
	if err != nil {
		return errors.NewInternalServerError("Error while adding the addreses")
	}
	return nil

}

func (u *userServices) RemoveAddress(userID string) {

}
