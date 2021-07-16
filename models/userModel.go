package models

type User struct {
	Email    string `json:"email" bson:"_id"`
	Password string `json:"password" bson:"password"`
}
