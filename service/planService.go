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
	ps, err := p.repo.GetListplanByUid(uid)
	res := []model.Plan{}
	for _, v := range ps {
		v.Key = MappingDBCate(v.Key)
		res = append(res, v)
	}
	return res, err
}

func (p *PlanService) Create(plan *model.Plan) error {
	plan.Id = uuid.NewString()
	plan.Timestamp = time.Now().UnixMilli()
	plan.Key = MappingCategory(plan.Key)
	err := p.repo.InsertPlan(plan)
	return err
}

func (p *PlanService) Update(plan *model.Plan) error {
	plan.Key = MappingCategory(plan.Key)
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
	key = MappingCategory(key)
	err := p.repo.DeletePlan(id, key)
	return err
}

func (p *PlanService) GetByKey(id string, key string) ([]model.Plan, error) {
	key = MappingCategory(key)
	pl, err := p.repo.GetPlanByKey(id, key)
	res := []model.Plan{}
	for _, v := range pl {
		v.Key = MappingDBCate(v.Key)
		res = append(res, v)
	}
	return res, err
}
