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
	headerXAccessTokenID = "X-Token-ID"
	headerXCallerID      = "X-Caller-ID"
)

var (
	headers         = make(http.Header)
	oauthRestClient = rest.RequestBuilder{
		BaseURL: "http://localhost:8090",
		Timeout: 200 * time.Millisecond,
		Headers: headers,
	}
)

type accessToken struct {
	UserID string `json:"user_id"`
	Email  string `json:"email"`
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

	accessTokenID := request.Header.Get(headerXAccessTokenID)
	if accessTokenID == "" {
		return errors.NewBadRequest("InValid Access Token")
	}
	userID := GetCallerID(request)
	if len(userID) <= 0 {
		return errors.NewBadRequest("Invalid User ID")
	}
	_, err := verifyAccessToken(accessTokenID, userID)
	if err != nil {
		if err.Status == http.StatusNotFound {
			return errors.NewNotFound("Token id is expired")
		}
		return err
	}
	return nil
}

func verifyAccessToken(accessTokenID, userID string) (*accessToken, *errors.RestError) {
	headers.Set(headerXCallerID, userID)
	headers.Set(headerXAccessTokenID, accessTokenID)
	response := oauthRestClient.Get("/validate")
	if response == nil || response.Response == nil {
		return nil, errors.NewNotFound("Not found")
	}

	if response.StatusCode > 299 {
		err := errors.NewInternalServerError("invalid access token id")
		return nil, err
	}

	var at accessToken
	if err := json.Unmarshal(response.Bytes(), &at); err != nil {
		fmt.Println("The value of error is", err)
		return nil, errors.NewInternalServerError("error when trying to unmarshal access token response")
	}
	return &at, nil
}
