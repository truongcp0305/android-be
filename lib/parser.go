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
			return nil, errors.New("Internal server error")
		}
		spends = append(spends, spend)
	}

	if err := cursor.Err(); err != nil {
		log.Printf("Error iterating over spending: %v\n", err)
		return nil, errors.New("Internal server error")
	}

	return spends, nil
}
