package model

type User struct {
	Name  string `json:"name" bson:"name"`
	Email string `json:"email" bson:"email"`
}
