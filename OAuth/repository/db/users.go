package db

import (
	"encoding/json"
	"time"

	"github.com/ankitanwar/GoAPIUtils/errors"
	"github.com/ankitanwar/userLoginWithOAuth/Oauth/domain/users"
	"github.com/mercadolibre/golang-restclient/rest"
)

var (
	restClinet = rest.RequestBuilder{
		BaseURL: "http://localhost:8081",
		Timeout: 100 * time.Millisecond,
	}
)

//UsersRepository : Interface for the repository for Rest User
type UsersRepository interface {
	Login(string, string) (*users.User, *errors.RestError)
}

type userRepository struct{}

//NewRestRepository : To create the Rest User Repository
func NewRestRepository() UsersRepository {
	return &userRepository{}
}

func (u *userRepository) Login(email, password string) (*users.User, *errors.RestError) {
	request := users.LoginRequest{
		Email:    email,
		Password: password,
	}
	response := restClinet.Post("/user/login", request)
	if response.Response == nil || response == nil {
		return nil, errors.NewNotFound("Invalid Rest Clinet Response")
	}
	if response.StatusCode > 299 {
		var restErr errors.RestError
		if err := json.Unmarshal(response.Bytes(), &restErr); err != nil {
			return nil, errors.NewInternalServerError("Error while Unmarshaling the error ")
		}
		return nil, &restErr
	}
	user := &users.User{}
	if err := json.Unmarshal(response.Bytes(), user); err != nil {
		return nil, errors.NewInternalServerError("Error While unmarshalling the user")
	}
	return user, nil
}
