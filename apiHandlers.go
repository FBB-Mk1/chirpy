package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type chirpValid struct {
	Valid bool `json:"valid"`
}

type chirpError struct {
	Error string `json:"error"`
}

type chirp struct {
	Body string `json:"body"`
}

func chirpValidateHandler(w http.ResponseWriter, r *http.Request) {
	// decode chirp
	decoder := json.NewDecoder(r.Body)
	c := chirp{}
	err := decoder.Decode(&c)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	// check if chirp is valid, if not respond with { "error": err }
	err = chirpChecker(c.Body)
	if err != nil {
		respBody := chirpError{Error: fmt.Sprint(err)}
		res, err := json.Marshal(respBody)
		if err != nil {
			w.WriteHeader(500)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		w.Write(res)
		return
	}
	// respond with { "valid": true }
	respBody := chirpValid{Valid: true}
	res, err := json.Marshal(respBody)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func chirpChecker(c string) error {
	if c == "" {
		return errors.New("no chirp found")
	}
	if len(c) < 5 {
		return errors.New("chirp too small")
	}
	if len(c) > 140 {
		return errors.New("chirp too long")
	}
	return nil
}
