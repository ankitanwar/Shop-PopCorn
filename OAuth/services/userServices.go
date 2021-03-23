package services

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/ankitanwar/GoAPIUtils/errors"
	mongo "github.com/ankitanwar/Shop-PopCorn/Oauth/database"
	"github.com/ankitanwar/Shop-PopCorn/Oauth/domain"
	"github.com/dgrijalva/jwt-go"
	"github.com/mercadolibre/golang-restclient/rest"
)

var (
	restClinet = rest.RequestBuilder{
		BaseURL: "http://localhost:8081",
		Timeout: 100 * time.Millisecond,
	}
	mySigningKey = []byte("AddThisKeyInEnvironmentVariable")
)

//CreateAccessToken : To create the new access Token
func CreateAccessToken(req *domain.LoginRequest) (*domain.AccessToken, *errors.RestError) {
	ans := restClinet.Post("/user/verify", req)
	if ans == nil || ans.Response == nil {
		return nil, errors.NewInternalServerError("Error while login")
	}
	user := &domain.User{}
	if ans.StatusCode < 299 {
		err := json.Unmarshal(ans.Bytes(), user)
		if err != nil {
			fmt.Println("The value of err is", err)
			return nil, errors.NewInternalServerError("Error while unmarshalling the data")
		}
	} else {
		return nil, errors.NewBadRequest("Unable to find the user")
	}
	tokenID := jwt.New(jwt.SigningMethodHS256)
	claims := tokenID.Claims.(jwt.MapClaims)
	claims["authorized"] = true
	claims["user"] = user.UserID
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()
	tokenString, err := tokenID.SignedString(mySigningKey)
	if err != nil {
		return nil, errors.NewInternalServerError("Error While Creating The Access Token")
	}
	token := &domain.AccessToken{
		UserID: user.UserID,
		Email:  user.Email,
		Token:  tokenString,
	}
	err = mongo.UpdateAccessToken(token)
	if err != nil {
		return nil, errors.NewInternalServerError("Unable To Create The Access Token")
	}
	return token, nil

}

//ValidateAccessToken : To validate the Access Token
func ValidateAccessToken(userID string, token string) (*domain.AccessToken, *errors.RestError) {
	check, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("There is an error")
		}
		return mySigningKey, nil
	})
	if err != nil {
		return nil, errors.NewInternalServerError("Error While Verify The Token")
	}
	if check.Valid {
		response, err := mongo.GetAccessToken(userID)
		if err != nil {
			return nil, errors.NewInternalServerError("Error WhileFethcing The Token")
		}
		if response.Token == "" {
			return nil, errors.NewBadRequest("Invalid Token ID")
		}
		return response, nil
	}
	return nil, errors.NewBadRequest("Invalid Token ID")
}

//RemoveAccessToken : To Remove The Access Token Of the Given User
func RemoveAccessToken(userID string, sentToken string) *errors.RestError {
	detail, err := mongo.GetAccessToken(userID)
	if err != nil {
		return errors.NewInternalServerError("Error While Fetching The details")
	}
	if detail.Token != sentToken {
		return errors.NewBadRequest("Invalid Access Token")
	}
	err = mongo.RemoveAccessToken(userID)
	if err != nil {
		return errors.NewInternalServerError("Error While removing The Access Token")
	}
	return nil

}
