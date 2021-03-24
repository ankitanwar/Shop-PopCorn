package user

import (
	"encoding/json"
	"net/http"
	"time"

	"fmt"

	"github.com/ankitanwar/GoAPIUtils/errors"
	"github.com/mercadolibre/golang-restclient/rest"
)

const (
	headerXAccessTokenID = "X-Token-ID"
	headerXCallerID      = "X-Caller-ID"
)

var (
	headers         = make(http.Header)
	oauthRestClient = rest.RequestBuilder{
		BaseURL: "http://localhost:8081",
		Timeout: 200 * time.Millisecond,
		Headers: headers,
	}
	//GetUserAddress : To get the address of the given user
	GetUserAddress userInterace = &user{}
)

type userInterace interface {
	GetAddress(*http.Request, string) (*SpecificAddress, *errors.RestError)
}
type user struct{}

//GetCallerID : To get the caller id from the url
func GetCallerID(request *http.Request) string {
	if request == nil {
		return ""
	}
	callerID := request.Header.Get(headerXCallerID)
	return callerID
}

//GetAccessID: To get the caller id from the url
func GetAccessID(request *http.Request) string {
	if request == nil {
		return ""
	}
	callerID := request.Header.Get(headerXAccessTokenID)
	return callerID
}

type SpecificAddress struct {
	ID          string
	HouseNumber string `json:"houseNo"`
	Street      string `json:"street"`
	State       string `json:"state"`
	Country     string `json:"country"`
	Phone       string `json:"phone"`
}

func (u *user) GetAddress(request *http.Request, addressID string) (*SpecificAddress, *errors.RestError) {
	userID := GetCallerID(request)
	accessTokenID := GetAccessID(request)
	headers.Set(headerXCallerID, userID)
	headers.Set(headerXAccessTokenID, accessTokenID)
	response := oauthRestClient.Get(fmt.Sprintf("/user/specificaddress/%s", addressID))
	if response == nil || response.Response == nil {
		return nil, errors.NewNotFound("Not found")
	}

	if response.StatusCode > 299 {
		err := errors.NewInternalServerError("Unable to fetch the address")
		return nil, err
	}

	address := &SpecificAddress{}
	if err := json.Unmarshal(response.Bytes(), &address); err != nil {
		return nil, errors.NewInternalServerError("error when trying to unmarshal access token response")
	}
	return address, nil

}
