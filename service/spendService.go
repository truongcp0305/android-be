package service

import (
	"android-be/model"
	"android-be/repository"
)

type SpendService struct {
	repo *repository.Database
}

func NewSpendService(rp *repository.Database) *SpendService {
	return &SpendService{
		repo: rp,
	}
}

func (s *SpendService) ListSpend(uid string) ([]model.Spending, error) {
	sps, err := s.repo.GetListSpendByUid(uid)
	return sps, err
}

func (s *SpendService) Create(spend *model.Spending) error {
	err := s.repo.InsertSpend(spend)
	return err
}

func (s *SpendService) UpdateSpend(spend *model.Spending) error {
	err := s.repo.UpdateSpend(spend)
	return err
}

func (s *SpendService) Delete(id string) error {
	err := s.repo.DeleteSpend(id)
	return err
}
