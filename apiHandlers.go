package main

import (
	"encoding/json"
	"net/http"
)

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
