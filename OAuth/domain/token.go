package domain

import (
	"time"

	"github.com/ankitanwar/GoAPIUtils/errors"
)

const (
	experationTime = 24
)

//AccessToken : Fields for accessToken
type AccessToken struct {
	UserID  int    `json:"user_id" bson:"user_id"`
	Email   string `json:"email" bson:"email"`
	Expires int64  `json:"expires" bson:"expires"`
	Token   string `json:"access_token" bson:"access_token"`
}

//ValidateAccessToken : To validate the access token
func ValidateAccessToken(token AccessToken) *errors.RestError {
	if token.UserID < 0 {
		return errors.NewBadRequest("Invalid User ID")
	}
	if len(token.Email) < 0 {
		return errors.NewBadRequest("Invalid Email Address")
	}
	exp := token.IsExpired()
	if exp == true {
		return errors.NewBadRequest("Access Token has been expired")
	}
	return nil

}

//GetNewAccessToken : To get the new access token
func (token *AccessToken) GetNewAccessToken() {
	token.Expires = time.Now().UTC().Add(experationTime * time.Hour).Unix()
}

//IsExpired : To Check whether the givenaccess token is experied or not
func (token *AccessToken) IsExpired() bool {
	return time.Unix(token.Expires, 0).Before(time.Now().UTC())
}
