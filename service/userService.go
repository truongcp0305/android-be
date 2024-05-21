package service

import (
	"android-be/model"
	"android-be/repository"

	"github.com/google/uuid"
)

type UserService struct {
	repo *repository.Database
}

func NewUserService(rp *repository.Database) *UserService {
	return &UserService{
		repo: rp,
	}
}

func (s *UserService) Login(u *model.User) (model.User, error) {
	return s.repo.Login(u)
}

func (s *UserService) CreateUser(u *model.User) (string, error) {
	u.Id = uuid.NewString()
	err := s.repo.CreateUser(u)
	if err != nil {
		return "", err
	}
	return u.Id, nil
}

func (s *UserService) GetInfo(id string) (model.User, error) {
	return s.repo.GetUserInfo(id)
}
