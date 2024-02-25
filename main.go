package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type apiConfig struct {
	fileserverHits int
}

func main() {
	apiCfg := apiConfig{}
	rt := chi.NewRouter()
	corsMux := middlewareCors(rt)
	fsHandler := apiCfg.middlewareMetricsInc(http.StripPrefix("/app", http.FileServer(http.Dir("."))))

	rt.Handle("/app", fsHandler)
	rt.Handle("/app/*", fsHandler)
	rt.Get("/healthz", healthzHandler)
	rt.Get("/metrics", apiCfg.metricsHandler)
	rt.HandleFunc("/reset", apiCfg.resetHandler)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: corsMux,
	}
	err := srv.ListenAndServe()

	if err != nil {
		fmt.Println(err)
	}
}
