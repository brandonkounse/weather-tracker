package client

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"weather-tracker/api"
	"weather-tracker/config"
)

func GetCurrentWeatherByLocation(city string) (*api.Current, error) {
	urlCurrentWeather := "http://api.weatherstack.com/current"

	apiKey, err := config.GetAPIKey()
	if err != nil {
		return &api.Current{}, err
	}

	u, err := url.Parse(urlCurrentWeather)
	if err != nil {
		log.Fatal("invalid base URL:", err)
	}

	q := u.Query()
	q.Set("access_key", apiKey)
	q.Set("query", city)
	u.RawQuery = q.Encode()

	reqURL := u.String()
	fmt.Println("Request URL:", reqURL)

	resp, err := http.Get(reqURL)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		b, _ := io.ReadAll(resp.Body)
		log.Fatalf("non-200 response: %d body =%s", resp.StatusCode, string(b))
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(bodyBytes))

	var apiCurrent api.Current

	if err := json.NewDecoder(resp.Body).Decode(&apiCurrent); err != nil {
		return &api.Current{}, err
	}

	return &apiCurrent, nil
}
