package repository

import (
	"android-be/lib"
	"android-be/model"
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Database struct {
	s *mongo.Collection
	p *mongo.Collection
	u *mongo.Collection
}

func NewDatabase(s *mongo.Collection, p *mongo.Collection, u *mongo.Collection) *Database {
	return &Database{
		s: s,
		p: p,
		u: u,
	}
}

func (d *Database) CreateUser(u *model.User) error {
	_, err := d.u.InsertOne(context.Background(), u)
	return err
}

func (d *Database) Login(user *model.User) (model.User, error) {
	filter := bson.M{"username": user.Username, "password": user.Password}
	c, err := d.u.Find(context.Background(), filter)
	if err != nil {
		return model.User{}, err
	}
	us, err := lib.ParseUser(c)
	if err != nil {
		return model.User{}, err
	}
	if len(us) == 0 {
		return model.User{}, errors.New("not found")
	}
	return us[0], nil
}

func (d *Database) GetUserInfo(id string) (model.User, error) {
	filter := bson.M{"id": id}
	c, err := d.u.Find(context.Background(), filter)
	if err != nil {
		return model.User{}, err
	}
	us, err := lib.ParseUser(c)
	if err != nil {
		return model.User{}, err
	}
	if len(us) == 0 {
		return model.User{}, errors.New("not found")
	}
	return us[0], nil
}

func (d *Database) GetSpendInWeek(time int64, uid string) ([]model.Spending, error) {
	filter := bson.M{
		"user_id":   uid,
		"timestamp": bson.M{"$gt": time},
	}
	c, err := d.s.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	return lib.ParseSpending(c)
}

func (d *Database) GetListSpendByUid(uid string) ([]model.Spending, error) {
	filter := bson.M{"user_id": uid}
	c, err := d.s.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	return lib.ParseSpending(c)
}

func (d *Database) GetSpend(id string) (model.Spending, error) {
	filter := bson.M{"id": id}
	c, err := d.s.Find(context.Background(), filter)
	if err != nil {
		return model.Spending{}, err
	}
	us, err := lib.ParseSpending(c)
	if err != nil {
		return model.Spending{}, err
	}
	if len(us) == 0 {
		return model.Spending{}, errors.New("not found")
	}
	return us[0], nil
}

func (d *Database) GetListplanByUid(uid string) ([]model.Plan, error) {
	filter := bson.M{"user_id": uid}
	c, err := d.p.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	return lib.ParsePlan(c)
}

func (d *Database) InsertPlan(plan *model.Plan) error {
	_, err := d.p.InsertOne(context.Background(), plan)
	return err
}

func (d *Database) InsertSpend(spend *model.Spending) error {
	_, err := d.s.InsertOne(context.Background(), spend)
	return err
}

func (d *Database) UpdateSpend(spend *model.Spending) error {
	filter := bson.M{"id": spend.Id}
	update := bson.M{"$set": spend}
	_, err := d.s.UpdateOne(context.TODO(), filter, update)
	return err
}

func (d *Database) DeleteSpend(id string) error {
	filter := bson.M{"id": id}
	_, err := d.s.DeleteOne(context.Background(), filter)
	return err
}

func (d *Database) UpdatePlan(plan *model.Plan) error {
	filter := bson.M{"user_id": plan.UserId, "key": plan.Key}
	update := bson.M{"$set": plan}
	_, err := d.p.UpdateOne(context.TODO(), filter, update)
	return err
}

func (d *Database) DeletePlan(id string, key string) error {
	filter := bson.M{"user_id": id, "key": key}
	_, err := d.p.DeleteOne(context.Background(), filter)
	return err
}

func (d *Database) GetSpendBetweenTime(uid string, start int64, end int64) ([]model.Spending, error) {
	filter := bson.M{
		"user_id":   uid,
		"timestamp": bson.M{"$gt": start, "$lt": end},
	}
	c, err := d.s.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	return lib.ParseSpending(c)
}

func (d *Database) GetPlanByKey(uid string, key string) ([]model.Plan, error) {
	filter := bson.M{"user_id": uid, "key": key}
	c, err := d.s.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	return lib.ParsePlan(c)
}
