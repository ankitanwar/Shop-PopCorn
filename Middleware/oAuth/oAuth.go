package oauth

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/ankitanwar/GoAPIUtils/errors"

	"github.com/mercadolibre/golang-restclient/rest"
)

const (
	headerXPublic   = "X-Public"
	headerXCallerID = "X-Caller-Id"

	paramAccessToken = "access_token"
)

var (
	oauthRestClient = rest.RequestBuilder{
		BaseURL: "http://localhost:8090",
		Timeout: 200 * time.Millisecond,
	}
)

type accessToken struct {
	UserID int64  `json:"user_id"`
	Email  string `json:"email"`
}

//IsPublic : To validate the request whether the request is public or not
func IsPublic(request *http.Request) bool {
	if request == nil {
		return true
	}
	return request.Header.Get(headerXPublic) == "true"
}

//GetCallerID : To get the caller id from the url
func GetCallerID(request *http.Request) string {
	if request == nil {
		return ""
	}
	callerID := request.Header.Get(headerXCallerID)
	return callerID
}

//AuthenticateRequest : To authenticate the given request
func AuthenticateRequest(request *http.Request) *errors.RestError {
	if request == nil {
		return errors.NewInternalServerError("Invalid Error")
	}

	accessTokenID := request.Header.Get(paramAccessToken)
	if accessTokenID == "" {
		return errors.NewBadRequest("InValid Access Token")
	}
	userID := GetCallerID(request)
	fmt.Println("The value of userId is ", userID)
	if len(userID) <= 0 {
		return errors.NewBadRequest("Invalid User ID")
	}
	at, err := getAccessToken(accessTokenID, userID)
	if err != nil {
		if err.Status == http.StatusNotFound {
			return errors.NewNotFound("Token id is expired")
		}
		return err
	}
	request.Header.Add(headerXCallerID, fmt.Sprintf("%v", at.UserID))
	return nil
}

func getAccessToken(accessTokenID, userID string) (*accessToken, *errors.RestError) {
	response := oauthRestClient.Get(fmt.Sprintf("/validate/%s/%s", userID, accessTokenID))
	if response == nil || response.Response == nil {
		return nil, errors.NewNotFound("Not found")
	}

	if response.StatusCode > 299 {
		err := errors.NewInternalServerError("invalid access token id")
		return nil, err
	}

	var at accessToken
	if err := json.Unmarshal(response.Bytes(), &at); err != nil {
		return nil, errors.NewInternalServerError("error when trying to unmarshal access token response")
	}
	return &at, nil
}
