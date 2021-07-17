package main

import (
	"fmt"
	"net/http"
	"os"

	"winkedIn/handlerFunctions"

	"github.com/gorilla/handlers"

	"github.com/gorilla/mux"
)

func GetPort() string {
	var port = os.Getenv("PORT")
	if port == "" {
		fmt.Println("INFO: No PORT environment variable detected, defaulting to 3000")
		return ":3000"
	}
	return ":" + port
}

func main() {
	r := mux.NewRouter()
	corsObj := handlers.AllowedOrigins([]string{"*"})
	r.HandleFunc("/register", handlerFunctions.Register).Methods("POST")
	r.HandleFunc("/profileRegister", handlerFunctions.Profile).Methods("POST")
	r.HandleFunc("/user/{email}", handlerFunctions.UserDetails).Methods("GET")
	r.HandleFunc("/explore/{email}", handlerFunctions.ExplorerUserDetails).Methods("GET")

	r.HandleFunc("/login", handlerFunctions.LoginUser).Methods("POST")

	r.HandleFunc("/api", api)
	http.ListenAndServe(GetPort(), handlers.CORS(corsObj)(r))
}
func api(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello API"))
	w.Write([]byte("\nPort: " + GetPort()))
}
