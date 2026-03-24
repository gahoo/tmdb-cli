package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/url"
)

// FindByExternalID looks up TV, Movie, etc. by external ID (e.g., imdb_id, tvdb_id)
func (c *Client) FindByExternalID(externalID string, externalSource string) (*FindResults, error) {
	endpoint := fmt.Sprintf("/find/%s?external_source=%s&language=en-US", url.PathEscape(externalID), externalSource)
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

	var result FindResults
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}
