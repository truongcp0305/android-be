package service

import (
	"android-be/model"
	"android-be/repository"
	"time"

	"github.com/google/uuid"
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
	plan.Id = uuid.NewString()
	plan.Timestamp = time.Now().UnixMilli()
	err := p.repo.InsertPlan(plan)
	return err
}

func (p *PlanService) Update(plan *model.Plan) error {
	err := p.repo.UpdatePlan(plan)
	return err
}

func (p *PlanService) Delete(id string, key string) error {

	// plans, err := p.GetByKey(id, key)
	// if err != nil {
	// 	return err
	// }

	// if len(plans) == 0 {
	// 	return nil
	// }
	err := p.repo.DeletePlan(id, key)
	return err
}

func (p *PlanService) GetByKey(id string, key string) ([]model.Plan, error) {
	return p.repo.GetPlanByKey(id, key)
}
