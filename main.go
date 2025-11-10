package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/joho/godotenv"
)

const (
	URL_CURRENT = "http://api.weatherstack.com/current"
)

// TODO standup a MVP, will refactor into proper files later
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		log.Fatal("API_KEY is empty")
	}

	u, err := url.Parse(URL_CURRENT)
	if err != nil {
		log.Fatal("invalid base URL:", err)
	}

	q := u.Query()
	q.Set("access_key", apiKey)
	q.Set("query", "Austin")
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
}
