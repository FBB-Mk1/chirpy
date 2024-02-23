package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	corsMux := middlewareCors(mux)
	mux.Handle("/app/", http.FileServer(http.Dir(".")))
	mux.HandleFunc("/healthz", healtHandler)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: corsMux,
	}
	err := srv.ListenAndServe()

	if err != nil {
		fmt.Println(err)
	}
}

func middlewareCors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func healtHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
