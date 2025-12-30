package service

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestGetWeatherByCEP(t *testing.T) {
	via := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{"localidade":"São Paulo"}`))
	}))
	defer via.Close()

	weather := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{"current":{"temp_c":25}}`))
	}))
	defer weather.Close()

	svc := NewWeatherService(via.URL, weather.URL, "key", 2*time.Second)

	resp, err := svc.GetWeatherByCEP(context.Background(), "01001000")
	if err != nil {
		t.Fatal(err)
	}
	if resp.TempC != 25 {
		t.Fatal("wrong temp")
	}
}
