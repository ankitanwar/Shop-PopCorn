package user

import (
	"encoding/json"
	"time"

	"fmt"

	"github.com/ankitanwar/GoAPIUtils/errors"
	"github.com/mercadolibre/golang-restclient/rest"
)

var (
	oauthRestClient = rest.RequestBuilder{
		BaseURL: "http://localhost:8081",
		Timeout: 200 * time.Millisecond,
	}
	//GetUserAddress : To get the address of the given user
	GetUserAddress userInterace = &user{}
)

type user struct{}

//Address : fields of address of the given user
type Address struct {
	UserID int             `json:"UserID"`
	List   []speficAddress `json:"List"`
}

type speficAddress struct {
	Address string `json:"address"`
	State   string `json:"state"`
	Country string `json:"country"`
	Phone   string `json:"phone"`
}

type userInterace interface {
	GetAddress(string) (*Address, *errors.RestError)
}

func (u *user) GetAddress(userID string) (*Address, *errors.RestError) {
	response := oauthRestClient.Get(fmt.Sprintf("/user/address/%s", userID))
	if response == nil || response.Response == nil {
		return nil, errors.NewNotFound("Not found")
	}

	if response.StatusCode > 299 {
		err := errors.NewInternalServerError("Unable to fetch the address")
		return nil, err
	}

	address := &Address{}
	if err := json.Unmarshal(response.Bytes(), &address); err != nil {
		return nil, errors.NewInternalServerError("error when trying to unmarshal access token response")
	}
	return address, nil

}
