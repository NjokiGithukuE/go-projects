package main

import (
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
)

func loginHandler(w http.ResponseWriter, r*http.Request) {
	var credentials struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := authenticate(credentials.Username, credentials.Password)
	if err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	jwtToken, err := generateJWT(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"token": jwtToken})
}

func secureResourceHandler(w http.ResponseWriter, r*http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("This is a secure resource accessible only to users with 'admin' role."))
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/login", loginHandler).Methods("POST")
	r.HandleFunc("/secure/resource", authorize("admin", secureResourceHandler)).Methods("GET")

	http.Handle("/", r)

	port := ":8080"
	println("Server is running on port", port)
	http.ListenAndServe(port, nil)
}