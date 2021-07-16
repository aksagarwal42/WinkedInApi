package handlerFunctions

import (
	"encoding/json"

	"net/http"
	"winkedIn/database"
	"winkedIn/models"

	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

// Register a User

func Register(w http.ResponseWriter, r *http.Request) {

	db, ctx := database.InitDB()

	userCollection := db.Collection("users")
	newUser := &models.User{
		Email:    r.FormValue("email"),
		Password: r.FormValue("password"),
	}

	json.NewDecoder(r.Body).Decode(&newUser)

	existingUser, _ := userCollection.CountDocuments(ctx, bson.D{{Key: "_id", Value: newUser.Email}})

	if existingUser != 0 {
		w.WriteHeader(http.StatusConflict)
		http.Error(w, "Email already exists", http.StatusConflict)
		return
	}

	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(newUser.Password), 14)
	newUser.Password = string(hashPassword)
	userCollection.InsertOne(ctx, newUser)

}

// Login a User

func LoginUser(w http.ResponseWriter, r *http.Request) {

	db, ctx := database.InitDB()

	registeredUsers := db.Collection("users")
	loginUser := &models.User{
		Email:    r.FormValue("email"),
		Password: r.FormValue("password"),
	}

	json.NewDecoder(r.Body).Decode(&loginUser)

	res := registeredUsers.FindOne(ctx, bson.D{{Key: "_id", Value: loginUser.Email}})

	if res.Err() != nil {
		w.WriteHeader(403)
		http.Error(w, "Invalid email", http.StatusForbidden)
		return
	}
	userDetails, _ := res.DecodeBytes()

	password := userDetails.Lookup("password").StringValue()
	err := bcrypt.CompareHashAndPassword([]byte(password), []byte(loginUser.Password))

	if err != nil {
		w.WriteHeader(403)
		http.Error(w, "Wrong Password", http.StatusForbidden)
		return
	}

	w.Write([]byte(loginUser.Email))
}
