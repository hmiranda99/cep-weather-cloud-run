package main

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/mateus/cep-weather-cloudrun/internal/service"
	"github.com/mateus/cep-weather-cloudrun/internal/validation"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	svc := service.NewWeatherService(
		"https://viacep.com.br",
		"https://api.weatherapi.com",
		os.Getenv("WEATHER_API_KEY"),
		10*time.Second,
	)

	http.HandleFunc("/weather", func(w http.ResponseWriter, r *http.Request) {
		cep := r.URL.Query().Get("cep")

		if !validation.IsValidCEP(cep) {
			w.WriteHeader(http.StatusUnprocessableEntity)
			w.Write([]byte("invalid zipcode"))
			return
		}

		resp, err := svc.GetWeatherByCEP(r.Context(), cep)
		if err != nil {
			if errors.Is(err, service.ErrZipNotFound) {
				w.WriteHeader(http.StatusNotFound)
				w.Write([]byte("can not find zipcode"))
				return
			}
			w.WriteHeader(http.StatusBadGateway)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	})

	log.Println("server running on port", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
