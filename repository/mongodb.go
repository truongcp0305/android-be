package repository

import (
	"android-be/model"

	"go.mongodb.org/mongo-driver/mongo"
)

type Database struct {
	s *mongo.Collection
	p *mongo.Collection
}

func (d *Database) GetListSpendByUid(uid string) ([]model.Spending, error) {
	return []model.Spending{}, nil
}

func (d *Database) InsertPlan(plan *model.Plan) error {
	return nil
}

func (d *Database) InsertSpend(spend *model.Spending) error {
	return nil
}

func (d *Database) UpdateSpend(spend *model.Spending) error {
	return nil
}
