package services

import (
	"github.com/ankitanwar/GoAPIUtils/errors"
	accesstoken "github.com/ankitanwar/userLoginWithOAuth/Oauth/domain/accessToken"
	"github.com/ankitanwar/userLoginWithOAuth/Oauth/repository/db"
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
	user, err := s.restUsersRepo.Login(request.Email, request.Password)
	if err != nil {
		return nil, err
	}
	token := accesstoken.GetNewAccessToken(user.ID)
	token.Generate()
	res, creatErr := s.repository.Create(token)
	if creatErr != nil {
		return nil, err
	}
	return res, nil
}

//GetById : To get the user buy the given id
func (s *service) GetByID(id string) (*accesstoken.AccessToken, *errors.RestError) {
	//id = strings.TrimSpace(id)
	if len(id) == 0 {
		return nil, errors.NewBadRequest("Invalid access token id")
	}
	token, err := s.repository.GetByID(id)
	if err != nil {
		return nil, errors.NewBadRequest("internal server error")
	}
	return token, nil

}

func (s *service) UpdateExperationTime(at accesstoken.AccessToken) *errors.RestError {
	if err := at.Validate(); err != nil {
		return errors.NewBadRequest("internal server error")
	}
	return errors.NewBadRequest("internal server error")
}
