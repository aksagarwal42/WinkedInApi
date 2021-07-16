package models

type Profile struct {
	Name       string   `json:"name" bson:"name"`
	DOB        string   `json:"dob" bson:"dob"`
	Image      string   `json:"image" bson:"image"`
	Email      string   `json:"email" bson:"email"`
	Gender     string   `json:"gender" bson:"gender"`
	City       string   `json:"city" bson:"city"`
	Bio        string   `json:"bio" bson:"bio"`
	Preference string   `json:"preference" bson:"preference"`
	Passion    string   `json:"passion" bson:"-"`
	Passions   []string `json:"passions" bson:"passions"`
}
