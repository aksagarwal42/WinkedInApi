package models

type UserProfile struct {
	Name     string   `json:"name" bson:"name"`
	City     string   `json:"city" bson:"city"`
	Gender   string   `json:"gender" bson:"gender"`
	Bio      string   `json:"bio" bson:"bio"`
	Passions []string `json:"passions" bson:"passions"`
	DOB      string   `json:"dob" bson:"dob"`
	Image    string   `json:"image" bson:"image"`
}
