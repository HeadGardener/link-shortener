package models

type User struct {
	ID       string `json:"id" bson:"id"`
	Name     string `json:"name" bson:"name"`
	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
}

type UserInput struct {
	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
}
