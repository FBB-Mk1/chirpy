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
	port := "8080"
	apiCfg := apiConfig{}
	rt := chi.NewRouter()
	apiRt := chi.NewRouter()
	corsMux := middlewareCors(rt)

	fsHandler := apiCfg.middlewareMetricsInc(http.StripPrefix("/app", http.FileServer(http.Dir("."))))

	rt.Mount("/api", apiRt)
	rt.Mount("/app", fsHandler)

	apiRt.Get("/healthz", healthzHandler)
	apiRt.Get("/metrics", apiCfg.metricsHandler)
	apiRt.HandleFunc("/reset", apiCfg.resetHandler)

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: corsMux,
	}

	err := srv.ListenAndServe()
	if err != nil {
		fmt.Println(err)
	}

}
