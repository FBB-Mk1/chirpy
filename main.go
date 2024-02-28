package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"www.github.com/fbb-mk1/chirpy/internal/database"
)

type apiConfig struct {
	fileserverHits int
}

func main() {
	port := "8080"
	apiCfg := apiConfig{}
	db, e := database.NewDB(".")
	if e != nil {
		log.Fatal("fuuuu")
	}
	db.CreateChirp("test")
	apiRt := chi.NewRouter()
	apiRt.Get("/healthz", healthzHandler)
	apiRt.Post("/validate_chirp", chirpValidateHandler)
	apiRt.HandleFunc("/reset", apiCfg.resetHandler)

	adminRt := chi.NewRouter()
	adminRt.Get("/metrics", apiCfg.metricsHandler)

	fsHandler := apiCfg.middlewareMetricsInc(http.StripPrefix("/app", http.FileServer(http.Dir("."))))
	rt := chi.NewRouter()

	rt.Mount("/api", apiRt)
	rt.Mount("/admin", adminRt)
	rt.Mount("/app", fsHandler)

	corsMux := middlewareCors(rt)

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: corsMux,
	}

	err := srv.ListenAndServe()
	if err != nil {
		fmt.Println(err)
	}

}
