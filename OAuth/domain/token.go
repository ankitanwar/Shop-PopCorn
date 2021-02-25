package domain

import (
	"fmt"
	"time"

	"github.com/ankitanwar/GoAPIUtils/errors"
	cryptos "github.com/ankitanwar/e-Commerce/User/utils/cryptoUtils"
)

const (
	experationTime = 24
)

//AccessToken : Fields for accessToken
type AccessToken struct {
	UserID  string `json:"user_id" bson:"_id"`
	Email   string `json:"email" bson:"email"`
	Expires int64  `json:"expires" bson:"expires"`
	Token   string `json:"access_token" bson:"access_token"`
}

//ValidateAccessToken : To validate the access token
func ValidateAccessToken(token AccessToken) *errors.RestError {
	if token.UserID == "" {
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

//CreateExperationTime : To get the new access token
func (token *AccessToken) CreateExperationTime() {
	token.Expires = time.Now().UTC().Add(experationTime * time.Hour).Unix()
}

//IsExpired : To Check whether the givenaccess token is experied or not
func (token *AccessToken) IsExpired() bool {
	return time.Unix(token.Expires, 0).Before(time.Now().UTC())
}

//CreateAccessToken : TO generate the new access token
func (token *AccessToken) CreateAccessToken() {
	token.Token = cryptos.GetMd5(fmt.Sprintf("at-%d-%d-ran", token.UserID, token.Expires))
}

//UpdateAccessToken : TO generate the new access token
func (token *AccessToken) UpdateAccessToken() string {
	return cryptos.GetMd5(fmt.Sprintf("at-%d-%d-ran", token.UserID, token.Expires))
}

//UpdateExperationTime : To get the new access token
func (token *AccessToken) UpdateExperationTime() int64 {
	return time.Now().UTC().Add(experationTime * time.Hour).Unix()
}
