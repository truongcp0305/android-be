package model

type Plan struct {
	Id        string `json:"id" bson:"id"`
	UserId    string `json:"user_id" bson:"user_id"`
	Key       string `json:"key" bson:"key"`
	Value     string `json:"value" bson:"value"`
	Plan      string `json:"plan" bson:"plan"`
	Timestamp int64  `json:"timestamp" bson:"timestamp"`
}
