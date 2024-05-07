package service

import (
	"android-be/model"
	"android-be/repository"
)

type PlanService struct {
	repo *repository.Database
}

func NewPalnService(rp *repository.Database) *PlanService {
	return &PlanService{
		repo: rp,
	}
}

func (p *PlanService) ListPlan(uid string) ([]model.Plan, error) {
	return p.repo.GetListplanByUid(uid)
}

func (p *PlanService) Create(plan *model.Plan) error {
	err := p.repo.InsertPlan(plan)
	return err
}

func (p *PlanService) Update(plan *model.Plan) error {
	err := p.repo.UpdatePlan(plan)
	return err
}

func (p *PlanService) Delete(id string) error {
	err := p.repo.DeletePlan(id)
	return err
}
