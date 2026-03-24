package api

import (
	"fmt"
	"net/http"
	"time"
)

const BaseURL = "https://api.themoviedb.org/3"

type Client struct {
	Token      string
	HTTPClient *http.Client
}

// NewClient creates a new TMDB API client instance with the given token
func NewClient(token string) *Client {
	return &Client{
		Token: token,
		HTTPClient: &http.Client{
			Timeout: 15 * time.Second,
		},
	}
}

// newRequest sets up a standard HTTP request to the TMDB API
func (c *Client) newRequest(method, endpoint string) (*http.Request, error) {
	url := fmt.Sprintf("%s%s", BaseURL, endpoint)
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}
	
	req.Header.Add("accept", "application/json")
	
	if c.Token != "" {
		// Suppress Bearer prefix logic for v3 keys vs v4 token if needed, 
		// but standard TMDB API uses Bearer for API Read Access Tokens string.
		req.Header.Add("Authorization", "Bearer "+c.Token)
	}
	
	return req, nil
}
