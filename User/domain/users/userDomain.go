package users

import (
	"strings"

	"github.com/ankitanwar/GoAPIUtils/errors"
	"github.com/uuid"
)

//User : User and its values
type User struct {
	ID          string `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
	Status      string `json:"status"`
	Password    string `json:"password"`
	PhoneNo     string `json:"phone"` //-> phone no is not neccessay while creating the acc
}

//Address : Address of the given user
type Address struct {
	UserID string        `bson:"_id" json:"userID"`
	List   []UserAddress `bson:"addresses" json:"addresses"`
}

//UserAddress : Address field for the user
type UserAddress struct {
	ID          string
	HouseNumber string `json:"houseNo"`
	Street      string `json:"street"`
	State       string `json:"state"`
	Country     string `json:"country"`
	Phone       string `json:"phone"`
}

//Users : It will return the slices of users
type Users []User

//ValidateAddress : To validate the given aaddress
func (address *UserAddress) ValidateAddress() *errors.RestError {
	if len(address.HouseNumber) < 0 {
		return errors.NewBadRequest("Enter the valid address")
	} else if len(address.Street) <= 0 {
		return errors.NewBadRequest("Please Enter The valid Street Number")
	} else if len(address.State) < 0 {
		return errors.NewBadRequest("Enter the valid address")
	} else if len(address.Country) < 0 {
		return errors.NewBadRequest("Enter the valid address")
	} else if len(address.Phone) > 10 || len(address.Phone) < 10 {
		return errors.NewBadRequest("Please Enter the valid phone number")
	}
	return nil
}

func (address *UserAddress) GenerateUniqueAddressID() (string, *errors.RestError) {
	id, err := uuid.NewV4()
	if err != nil {
		return "", errors.NewInternalServerError("Error While Generating The Address")
	}
	stringID := id.String()
	return stringID, nil
}

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
	if user.Password == "" || len(user.Password) < 5 {
		return errors.NewBadRequest("Please Enter the valid password")
	}
	if len(user.PhoneNo) != 0 {
		if len(user.PhoneNo) > 10 || len(user.PhoneNo) < 10 {
			return errors.NewBadRequest("Please Enter the valid phone Number")
		}
	}
	return nil
}
