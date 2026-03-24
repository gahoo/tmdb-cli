package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/url"
)

// Search finds movies, tv shows, or multiple types depending on searchType ("movie", "tv", "multi")
func (c *Client) Search(query string, searchType string) (*SearchResultPage, error) {
	if searchType != "movie" && searchType != "tv" && searchType != "multi" {
		return nil, fmt.Errorf("invalid search type: %s", searchType)
	}

	endpoint := fmt.Sprintf("/search/%s?query=%s&include_adult=false&language=en-US&page=1", searchType, url.QueryEscape(query))
	
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
