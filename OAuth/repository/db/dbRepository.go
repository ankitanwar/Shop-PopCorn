package db

import (
	mongodb "github.com/ankitanwar/OAuth/clients"
	accesstoken "github.com/ankitanwar/OAuth/domain/accessToken"
	"github.com/ankitanwar/OAuth/utils/errors"
	"go.mongodb.org/mongo-driver/bson"
)

var collections = mongodb.Client.Database("learning").Collection("people")

//Repository : Database Interface
type Repository interface {
	GetByID(string) (*accesstoken.AccessToken, *errors.RestError)
	Create(accesstoken.AccessToken) (*accesstoken.AccessToken, *errors.RestError)
	UpdateExperationTime(accesstoken.AccessToken) *errors.RestError
}

type dbRepository struct {
}

//NewRepository : It will return the pointer to the dbRepository interface
func NewRepository() Repository {
	return &dbRepository{}
}

func (d *dbRepository) GetByID(ID string) (*accesstoken.AccessToken, *errors.RestError) {
	ctx, cancel := mongodb.GetSession()
	defer cancel()
	var result accesstoken.AccessToken
	err := collections.FindOne(ctx, bson.M{"access_token": ID}).Decode(&result)
	if err != nil {
		return nil, errors.NewNotFound("Invalid Access Token")
	}
	return &result, nil
}

func (d *dbRepository) Create(at accesstoken.AccessToken) (*accesstoken.AccessToken, *errors.RestError) {
	session, close := mongodb.GetSession()
	defer close()
	_, err := collections.InsertOne(session, at)
	if err != nil {
		return nil, errors.NewInternalServerError("Error while getting the access token")
	}
	return &at, nil

}

func (d *dbRepository) UpdateExperationTime(at accesstoken.AccessToken) *errors.RestError {
	return nil

}
