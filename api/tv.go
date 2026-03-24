package api

import (
	"encoding/json"
	"fmt"
	"io"
)

// GetTVDetails retrieves the tv show details by ID
func (c *Client) GetTVDetails(tvID int) (*TVDetails, error) {
	endpoint := fmt.Sprintf("/tv/%d?language=en-US", tvID)
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

	var result TVDetails
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}

// GetTVSeasonDetails retrieves the tv season details by tv ID and season number
func (c *Client) GetTVSeasonDetails(tvID int, seasonNumber int) (*TVSeason, error) {
	endpoint := fmt.Sprintf("/tv/%d/season/%d?language=en-US", tvID, seasonNumber)
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

	var result TVSeason
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}
