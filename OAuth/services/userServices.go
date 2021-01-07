package services

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/ankitanwar/GoAPIUtils/errors"
	mongo "github.com/ankitanwar/e-Commerce/Oauth/database"
	"github.com/ankitanwar/e-Commerce/Oauth/domain"
	"github.com/mercadolibre/golang-restclient/rest"
	"go.mongodb.org/mongo-driver/bson"
)

var (
	restClinet = rest.RequestBuilder{
		BaseURL: "http://localhost:8081",
		Timeout: 100 * time.Millisecond,
	}
	//NotFound : If the id is not present in the database
	NotFound = "mongo: no documents in result"
)

//CreateAccessToken : To create the new access Token
func CreateAccessToken(req *domain.LoginRequest) (*domain.AccessToken, *errors.RestError) {
	ans := restClinet.Post("/user/login", req)
	if ans == nil || ans.Response == nil {
		return nil, errors.NewInternalServerError("Error while login")
	}
	user := &domain.User{}
	if ans.StatusCode < 299 {
		err := json.Unmarshal(ans.Bytes(), user)
		if err != nil {
			return nil, errors.NewInternalServerError("Error while unmarshalling the data")
		}
	}
	fmt.Println("The value of domain is ", user)
	token := &domain.AccessToken{}
	filter := bson.M{"user_id": user.UserID}
	findErr := mongo.Collection.FindOne(context.Background(), filter).Decode(token)
	if findErr != nil {
		if findErr.Error() == NotFound {
			token.UserID = user.UserID
			token.Email = user.Email
		}
		return nil, errors.NewInternalServerError("Cannot find the user")
	}
	token.GetNewAccessToken()
	_, updateErr := mongo.Collection.UpdateOne(context.Background(), filter, token)
	if updateErr != nil {
		return nil, errors.NewInternalServerError("Erro while updating the access token")
	}
	return token, nil

}

//ValidateAccessToken : To validate the Access Token
func ValidateAccessToken(id int, token string) (*domain.AccessToken, *errors.RestError) {
	filter := bson.M{"user_id": id}
	access := &domain.AccessToken{}
	err := mongo.Collection.FindOne(context.Background(), filter).Decode(access)
	if err != nil {
		return nil, errors.NewInternalServerError("User not found")
	}
	if access.Token == token {
		err := domain.ValidateAccessToken(*access)
		if err != nil {
			return nil, err
		}
		return access, nil
	}
	return nil, errors.NewBadRequest("Invalid Credentials")

}
