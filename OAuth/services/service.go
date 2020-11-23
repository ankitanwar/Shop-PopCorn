package services

import (
	accesstoken "github.com/ankitanwar/OAuth/domain/accessToken"
	"github.com/ankitanwar/OAuth/repository/db"
	"github.com/ankitanwar/OAuth/utils/errors"
)

//Service : Interface for services
type Service interface {
	GetByID(string) (*accesstoken.AccessToken, *errors.RestError)
	Create(*accesstoken.TokenRequest) (*accesstoken.AccessToken, *errors.RestError)
	UpdateExperationTime(accesstoken.AccessToken) *errors.RestError
}

type service struct {
	restUsersRepo db.UsersRepository
	repository    db.Repository
}

//NewService : it will return the pointer to the Service interface
func NewService(repo db.Repository, usersRepo db.UsersRepository) Service {
	return &service{
		restUsersRepo: usersRepo,
		repository:    repo,
	}
}

func (s *service) Create(request *accesstoken.TokenRequest) (*accesstoken.AccessToken, *errors.RestError) {
	user, err := s.restUsersRepo.Login(request.Email,request.Password)
	if err != nil {
		return nil, err
	}
	token := accesstoken.GetNewAccessToken(user.ID)
	token.Generate()
	_,err   = s.repository.Create(*token)
	if err != nil {
		return nil, err
	}
	return token, nil
}

//GetById : To get the user by the given id
func (s *service) GetByID(id string) (*accesstoken.AccessToken, *errors.RestError) {
	if len(id) == 0 {
		return nil, errors.NewBadRequest("Invalid access token id")
	}
	token, err := s.repository.GetByID(id)
	if err != nil {
		return nil, err
	}
	return token, nil

}

func (s *service) UpdateExperationTime(at accesstoken.AccessToken) *errors.RestError {
	if err := at.Validate(); err != nil {
		return err
	}
	return s.repository.UpdateExperationTime(at)
}
