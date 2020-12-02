package db

import (
	mongod "github.com/ankitanwar/e-Commerce/Oauth/clients"
	accesstoken "github.com/ankitanwar/e-Commerce/Oauth/domain/accessToken"
	"go.mongodb.org/mongo-driver/bson"

	"github.com/ankitanwar/GoAPIUtils/errors"
)

var collections = mongod.Client.Database("Users").Collection("user")

//Repository : Database Interface
type Repository interface {
	GetByID(string) (*accesstoken.AccessToken, *errors.RestError)
	Create(*accesstoken.AccessToken) (*accesstoken.AccessToken, *errors.RestError)
	UpdateExperationTime(accesstoken.AccessToken) *errors.RestError
}

type dbRepository struct {
}

//NewRepository : It will return the pointer to the dbRepository interface
func NewRepository() Repository {
	return &dbRepository{}
}

func (d *dbRepository) GetByID(ID string) (*accesstoken.AccessToken, *errors.RestError) {
	ctx, cancel := mongod.GetSession()
	defer cancel()
	result := &accesstoken.AccessToken{}
	err := collections.FindOne(ctx, bson.M{"access_token": ID}).Decode(result)
	if err != nil {
		return nil, errors.NewNotFound("Given ID doesnt found in the database")
	}
	return result, nil
}

func (d *dbRepository) Create(at *accesstoken.AccessToken) (*accesstoken.AccessToken, *errors.RestError) {
	session, close := mongod.GetSession()
	defer close()
	_, err := collections.InsertOne(session, at)
	if err != nil {
		return nil, errors.NewInternalServerError("Error while getting the access token")
	}
	return at, nil

}

func (d *dbRepository) UpdateExperationTime(at accesstoken.AccessToken) *errors.RestError {
	return nil

}
