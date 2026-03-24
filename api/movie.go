package api

import (
	"encoding/json"
	"fmt"
	"io"
)

// GetMovieDetails retrieves the movie details by ID
func (c *Client) GetMovieDetails(movieID int) (*MovieDetails, error) {
	endpoint := fmt.Sprintf("/movie/%d?language=en-US", movieID)
	req, err := c.newRequest("GET", endpoint)
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

	var result MovieDetails
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}
