package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type user struct {
	Email string `json:"email"`
}

func (cfg *apiConfig) userValidateHandler(w http.ResponseWriter, r *http.Request) {
	// decode user
	decoder := json.NewDecoder(r.Body)
	u := user{}
	err := decoder.Decode(&u)
	if err != nil {
		respondWithError(w, 500, err.Error())
		return
	}
	// check if user is valid, if not respond with { "error": err }
	ok := emailValidator(u.Email)
	if !ok {
		respondWithError(w, 400, "invalid email")
		return
	}
	user, err := cfg.db.CreateUser(u.Email)
	if err != nil {
		fmt.Println(err)
		respondWithJSON(w, 400, user)
		return
	}
	respondWithJSON(w, 201, user)
}

func emailValidator(email string) bool {
	//check if email is valid
	return true
}
