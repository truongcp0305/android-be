package model

type Plan struct {
	Id    string `json:"id" bson:"id"`
	Key   string `json:"key" bson:"key"`
	Value string `json:"value" bson:"value"`
	Plan  string `json:"plan" bson:"plan"`
}
