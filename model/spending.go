package model

type Spending struct {
	Id       string `json:"id" bson:"id"`
	Money    string `json:"money" bson:"money"`
	Icon     string `json:"icon" bson:"icon"`
	Category string `json:"category" bson:"category"`
	Time     string `json:"time" bson:"time"`
	Type     string `json:"type" bson:"type"`
	Note     string `json:"note" bson:"note"`
}
