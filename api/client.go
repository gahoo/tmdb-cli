package api

import (
	"fmt"
	"net/http"
	"time"
)

const BaseURL = "https://api.themoviedb.org/3"

type Client struct {
	Token      string
	Language   string
	HTTPClient *http.Client
}

// NewClient creates a new TMDB API client instance with the given token
func NewClient(token string, language string) *Client {
	return &Client{
		Token:    token,
		Language: language,
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
	
	q := req.URL.Query()
	if c.Language != "" {
		q.Add("language", c.Language)
	}
	req.URL.RawQuery = q.Encode()

	req.Header.Add("accept", "application/json")
	
	if c.Token != "" {
		req.Header.Add("Authorization", "Bearer "+c.Token)
	}
	
	return req, nil
}
