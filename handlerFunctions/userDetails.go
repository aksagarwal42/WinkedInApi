package handlerFunctions

import (
	"encoding/json"
	"net/http"
	"winkedIn/database"
	"winkedIn/models"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func UserDetails(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	db, ctx := database.InitDB()

	usersCollection := db.Collection("users")
	profileCollection := db.Collection("profile")

	var user models.UserProfile

	usersCollection.FindOne(ctx, bson.M{"_id": vars["email"]}).Decode(&user)
	profileCollection.FindOne(ctx, bson.M{"email": vars["email"]}).Decode(&user)
	user.DOB = user.DOB[6:]

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(user)
}

func ExplorerUserDetails(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	db, ctx := database.InitDB()

	profileCollection := db.Collection("profile")

	var user []models.Explorer
	var preference models.Preferences
	var filter primitive.D

	// Finding the  preference of the logged in user
	profileCollection.FindOne(ctx, bson.M{"email": vars["email"]}).Decode(&preference)

	if preference.Preference != "Both" {
		filter = bson.D{{Key: "gender", Value: preference.Preference}, {Key: "email", Value: bson.D{{Key: "$ne", Value: vars["email"]}}}}
	} else {
		filter = bson.D{{Key: "email", Value: bson.D{{Key: "$ne", Value: vars["email"]}}}}
	}

	// Finding possible users with gender = preference of the logged in user
	profileCursor, _ := profileCollection.Find(ctx, filter)

	// Decoding all users and storing it in user struct
	profileCursor.All(ctx, &user)

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(user)
}
