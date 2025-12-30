package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

var ErrZipNotFound = errors.New("zipcode not found")

type WeatherService struct {
	viaCEPBase  string
	weatherBase string
	apiKey      string
	httpClient  *http.Client
}

type Response struct {
	TempC float64 `json:"temp_C"`
	TempF float64 `json:"temp_F"`
	TempK float64 `json:"temp_K"`
}

func NewWeatherService(via, weather, key string, timeout time.Duration) *WeatherService {
	return &WeatherService{
		viaCEPBase:  via,
		weatherBase: weather,
		apiKey:      key,
		httpClient:  &http.Client{Timeout: timeout},
	}
}

func (s *WeatherService) GetWeatherByCEP(ctx context.Context, cep string) (Response, error) {
	city, err := s.lookupCity(ctx, cep)
	if err != nil {
		return Response{}, err
	}

	tempC, err := s.lookupTemp(ctx, city)
	if err != nil {
		return Response{}, err
	}

	return Response{
		TempC: tempC,
		TempF: tempC*1.8 + 32,
		TempK: tempC + 273,
	}, nil
}

func (s *WeatherService) lookupCity(ctx context.Context, cep string) (string, error) {
	u := fmt.Sprintf("%s/ws/%s/json/", s.viaCEPBase, cep)
	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, u, nil)
	resp, err := s.httpClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var payload struct {
		Localidade string `json:"localidade"`
		Erro       bool   `json:"erro"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&payload); err != nil {
		return "", err
	}

	if payload.Erro || payload.Localidade == "" {
		return "", ErrZipNotFound
	}

	return payload.Localidade, nil
}

func (s *WeatherService) lookupTemp(ctx context.Context, city string) (float64, error) {
	if s.apiKey == "" {
		return 0, errors.New("missing api key")
	}
	q := url.QueryEscape(city)
	u := fmt.Sprintf("%s/v1/current.json?key=%s&q=%s&aqi=no", s.weatherBase, s.apiKey, q)

	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, u, nil)
	resp, err := s.httpClient.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	var payload struct {
		Current struct {
			TempC float64 `json:"temp_c"`
		} `json:"current"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&payload); err != nil {
		return 0, err
	}

	return payload.Current.TempC, nil
}
