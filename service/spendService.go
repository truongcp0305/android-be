package service

import (
	"android-be/model"
	"android-be/repository"
	"time"

	"github.com/google/uuid"
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
	spend.Id = uuid.NewString()
	spend.Timestamp = time.Now().UnixMilli()
	err := s.repo.InsertSpend(spend)
	return err
}

func (s *SpendService) GetBill(id string) (model.Spending, error) {
	return s.repo.GetSpend(id)
}

func (s *SpendService) UpdateSpend(spend *model.Spending) error {
	err := s.repo.UpdateSpend(spend)
	return err
}

func (s *SpendService) Delete(id string) error {
	err := s.repo.DeleteSpend(id)
	return err
}

func (s *SpendService) GetInWeek(uid string) ([]model.Spending, error) {
	var part = time.Now().AddDate(0, 0, -7).UnixMilli()
	return s.repo.GetSpendInWeek(part, uid)
}
