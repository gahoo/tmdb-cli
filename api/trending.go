package api

import (
	"encoding/json"
	"fmt"
	"io"
)

// GetTrending retrieves trending items (all, movie, tv, person) over a time window (day, week)
func (c *Client) GetTrending(mediaType string, timeWindow string) (*SearchResultPage, error) {
	if mediaType != "all" && mediaType != "movie" && mediaType != "tv" && mediaType != "person" {
		return nil, fmt.Errorf("invalid media_type: %s", mediaType)
	}
	if timeWindow != "day" && timeWindow != "week" {
		return nil, fmt.Errorf("invalid time_window: %s", timeWindow)
	}

	endpoint := fmt.Sprintf("/trending/%s/%s?language=en-US", mediaType, timeWindow)
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

	var result SearchResultPage
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}
