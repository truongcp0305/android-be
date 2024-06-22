package model

type User struct {
	Id            string `json:"id" bson:"id"`
	Username      string `json:"username" bson:"username"`
	LastActive    string `json:"last_active" bson:"lastActive"`
	DeleteExpired string `json:"delete_expired" bson:"deleteExpired"`
	Password      string `json:"password" bson:"password"`
}
