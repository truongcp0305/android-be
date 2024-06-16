package service

import (
	"android-be/model"
	"android-be/repository"
	"fmt"
	"time"

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
	user, err := s.repo.Login(u)
	if err != nil {
		return user, err
	}
	user.LastActive = fmt.Sprintf("%d", time.Now().UnixMilli())
	user.DeleteExpired = ""
	err = s.repo.UpdateUser(&user)
	if err != nil {
		return user, err
	}
	return user, nil
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

func (s *UserService) AdLogin(ad *model.AdModel) error {
	if err := s.repo.AdminLogin(ad); err != nil {
		return err
	}
	return nil
}

func (s *UserService) ChangeUserInfo(u *model.User) error {
	if err := s.repo.UpdateUser(u); err != nil {
		return err
	}
	return nil
}

func (s *UserService) Search(data string) ([]model.User, error) {
	us, err := s.repo.Search(data)
	if err != nil {
		return us, err
	}
	return us, nil
}

func (s *UserService) QueryUser(page int) ([]model.User, error) {
	us, err := s.repo.QueryUser(page)
	if err != nil {
		return us, err
	}
	return us, nil
}
