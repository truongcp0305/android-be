package service

import "android-be/repository"

type SpendService struct {
	repo repository.Database
}

func NewSpendService(rp repository.Database) *SpendService {
	return &SpendService{
		repo: rp,
	}
}

func (s *SpendService) ListSpend(uid string) {

}
