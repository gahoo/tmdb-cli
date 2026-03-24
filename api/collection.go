package api

import (
	"encoding/json"
	"fmt"
	"io"
)

// GetCollection retrieves collection details by ID
func (c *Client) GetCollection(collectionID int) (*Collection, error) {
	endpoint := fmt.Sprintf("/collection/%d?language=en-US", collectionID)
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

	var result Collection
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}
