package api

import (
	"encoding/json"
	"fmt"
	"io"
)

// GetConfiguration retrieves the system wide configuration information for TMDB API
func (c *Client) GetConfiguration() (*Configuration, error) {
	req, err := c.newRequest("GET", "/configuration")
	if err != nil {
		return nil, err
	}

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		body, _ := io.ReadAll(res.Body)
		return nil, fmt.Errorf("API error: status code %d, body: %s", res.StatusCode, string(body))
	}

	var result Configuration
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}
