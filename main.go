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
	db             *database.DB
}

func main() {
	port := "8080"
	dbPath := "."

	db, e := database.NewDB(dbPath)
	if e != nil {
		log.Fatal("fuuuu")
	}
	apiCfg := apiConfig{fileserverHits: 0, db: db}

	apiRt := chi.NewRouter()
	apiRt.Get("/healthz", healthzHandler)
	apiRt.Post("/chirps", apiCfg.chirpValidateHandler)
	apiRt.Get("/chirps", apiCfg.getChirpHandler)
	apiRt.Get("/chirps/{id}", apiCfg.getSingleChirp)
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
