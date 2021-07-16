package models

type Explorer struct {
	Image string `json:"image" bson:"image"`
	Name  string `json:"name" bson:"name"`
	City  string `json:"city" bson:"city"`
	DOB   string `json:"dob" bson:"dob"`
}
