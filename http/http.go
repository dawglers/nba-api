package http

import (
	"fmt"
	"net/http"
	"net/url"
)

var headers = map[string]string{
	"Host":               "stats.nba.com",
	"User-Agent":         "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:72.0) Gecko/20100101 Firefox/72.0",
	"Accept":             "application/json",
	"Accept-Language":    "en-US,en;q=0.9,la;q=0.8",
	"Accept-Encoding":    "identity",
	"x-nba-stats-origin": "stats",
	"x-nba-stats-token":  "true",
	"Connection":         "keep-alive",
	"Referer":            "https://stats.nba.com/",
	"Pragma":             "no-cache",
	"Cache-Control":      "no-cache",
}

func MakeRequest(url *url.URL) (*http.Response, error) {
	// Build an HTTP request
	request, err := buildRequest(url)
	if err != nil {
		return nil, err
	}

	// Create a new HTTP client
	tr := &http.Transport{
		DisableCompression: true,
	}
	client := http.Client{Transport: tr}

	// Send the request via the HTTP client
	response, err := client.Do(request)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}

	if response.StatusCode != 200 {
		return nil, fmt.Errorf("error code received: %s", response.Status)
	}

	return response, nil
}

func buildRequest(url *url.URL) (*http.Request, error) {
	request, err := http.NewRequest("GET", url.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	for key, value := range headers {
		request.Header.Add(key, value)
	}

	return request, nil
}
