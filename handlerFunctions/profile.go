package handlerFunctions

import (
	"encoding/json"
	"net/http"
	"strings"
	"winkedIn/database"
	"winkedIn/models"
)

func Profile(w http.ResponseWriter, r *http.Request) {

	db, ctx := database.InitDB()

	profileCollection := db.Collection("profile")

	newUser := &models.Profile{
		Name:       r.FormValue("name"),
		DOB:        r.FormValue("dob"),
		Email:      r.FormValue("email"),
		City:       r.FormValue("city"),
		Gender:     r.FormValue("gender"),
		Bio:        r.FormValue("bio"),
		Passion:    r.FormValue("passion"),
		Image:      r.FormValue("image"),
		Preference: r.FormValue("preference"),
	}

	json.NewDecoder(r.Body).Decode(&newUser)

	newUser.Passions = strings.Split(newUser.Passion, ",")
	profileCollection.InsertOne(ctx, newUser)

}
