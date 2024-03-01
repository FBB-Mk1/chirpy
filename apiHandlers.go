package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

type chirpError struct {
	Error string `json:"error"`
}

type chirp struct {
	Body string `json:"body"`
}

func (cfg *apiConfig) getChirpHandler(w http.ResponseWriter, r *http.Request) {
	respBody, err := cfg.db.GetChirps()
	if err != nil {
		respondWithError(w, 500, err.Error())
	}
	respondWithJSON(w, 201, respBody)
}

func (cfg *apiConfig) chirpValidateHandler(w http.ResponseWriter, r *http.Request) {
	// decode chirp
	decoder := json.NewDecoder(r.Body)
	c := chirp{}
	err := decoder.Decode(&c)
	if err != nil {
		respondWithError(w, 500, err.Error())
		return
	}
	// check if chirp is valid, if not respond with { "error": err }
	err = chirpChecker(c.Body)
	if err != nil {
		respBody := chirpError{Error: fmt.Sprint(err)}
		respondWithJSON(w, 400, respBody)
		return
	}
	// respond with { "valid": true }
	body := c.Body
	cleaned, ok := checkProfanity(body)
	if !ok {
		body = cleaned
	}
	chirp, err := cfg.db.CreateChirp(body)
	if err != nil {
		respondWithJSON(w, 400, chirp)
		return
	}
	respondWithJSON(w, 201, chirp)
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(code)
	w.Write([]byte(msg))
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	res, err := json.Marshal(payload)
	if err != nil {
		code = 500
	}
	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

func chirpChecker(s string) error {
	if s == "" {
		return errors.New("no chirp found")
	}
	if len(s) > 140 {
		return errors.New("chirp too long")
	}
	return nil
}

func checkProfanity(s string) (string, bool) {
	profane := []string{"kerfuffle", "sharbert", "fornax"}
	words := strings.Split(s, " ")
	clean := true
	for i, w := range words {
		for _, p := range profane {
			if strings.ToLower(w) == p {
				words[i] = "****"
				clean = false
			}
		}
	}
	return strings.Join(words, " "), clean
}
