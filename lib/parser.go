package lib

import (
	"android-be/model"
	"context"
	"errors"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
)

func ParseSpending(cursor *mongo.Cursor) ([]model.Spending, error) {
	var spends []model.Spending
	defer cursor.Close(context.TODO())
	for cursor.Next(context.TODO()) {
		var spend model.Spending
		err := cursor.Decode(&spend)
		if err != nil {
			log.Printf("Error decoding spending: %v\n", err)
			return nil, errors.New("internal server error")
		}
		spends = append(spends, spend)
	}

	if err := cursor.Err(); err != nil {
		log.Printf("Error iterating over spending: %v\n", err)
		return nil, errors.New("internal server error")
	}

	return spends, nil
}

func ParsePlan(cursor *mongo.Cursor) ([]model.Plan, error) {
	var plans []model.Plan
	defer cursor.Close(context.TODO())
	for cursor.Next(context.TODO()) {
		var plan model.Plan
		err := cursor.Decode(&plan)
		if err != nil {
			log.Printf("Error decoding plan: %v\n", err)
			return nil, errors.New("internal server error")
		}
		plans = append(plans, plan)
	}

	if err := cursor.Err(); err != nil {
		log.Printf("Error iterating over spending: %v\n", err)
		return nil, errors.New("internal server error")
	}

	return plans, nil
}

func ParseUser(cursor *mongo.Cursor) ([]model.User, error) {
	var users []model.User
	defer cursor.Close(context.TODO())
	for cursor.Next(context.TODO()) {
		var user model.User
		err := cursor.Decode(&user)
		if err != nil {
			log.Printf("Error decoding plan: %v\n", err)
			return nil, errors.New("internal server error")
		}
		users = append(users, user)
	}

	if err := cursor.Err(); err != nil {
		log.Printf("Error iterating over spending: %v\n", err)
		return nil, errors.New("internal server error")
	}

	return users, nil
}
