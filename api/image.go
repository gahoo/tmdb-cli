package api

import (
	"fmt"
	"io"
	"os"
)

// DownloadImage retrieves an image from TMDB and saves it to a local path
func (c *Client) DownloadImage(posterPath string, destPath string) error {
	if posterPath == "" {
		return nil // Nothing to download
	}

	config, err := c.GetConfiguration()
	if err != nil {
		return fmt.Errorf("failed to fetch API configuration: %v", err)
	}

	baseURL := config.Images.SecureBaseURL
	if baseURL == "" {
		baseURL = "https://image.tmdb.org/t/p/"
	}

	url := fmt.Sprintf("%soriginal%s", baseURL, posterPath)
	resp, err := c.HTTPClient.Get(url)
	if err != nil {
		return fmt.Errorf("failed to download image: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return fmt.Errorf("failed to download image, status code: %d", resp.StatusCode)
	}

	file, err := os.Create(destPath)
	if err != nil {
		return fmt.Errorf("failed to create local file: %v", err)
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return fmt.Errorf("failed to write image to file: %v", err)
	}

	return nil
}
