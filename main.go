package main

import (
	"net/http"

	"winkedIn/handlerFunctions"

	"github.com/gorilla/handlers"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	corsObj := handlers.AllowedOrigins([]string{"*"})
	r.HandleFunc("/register", handlerFunctions.Register).Methods("POST")
	r.HandleFunc("/profileRegister", handlerFunctions.Profile).Methods("POST")
	r.HandleFunc("/user/{email}", handlerFunctions.UserDetails).Methods("GET")
	r.HandleFunc("/explore/{email}", handlerFunctions.ExplorerUserDetails).Methods("GET")

	r.HandleFunc("/login", handlerFunctions.LoginUser).Methods("POST")

	r.HandleFunc("/api", api)
	http.ListenAndServe("winkedinapi.herokuapp.com", handlers.CORS(corsObj)(r))
}
func api(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello API"))
}
