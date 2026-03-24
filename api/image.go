package api

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

// DownloadImage retrieves an image from TMDB and saves it to a local path
func DownloadImage(posterPath string, destPath string) error {
	if posterPath == "" {
		return nil // Nothing to download
	}

	url := fmt.Sprintf("https://image.tmdb.org/t/p/original%s", posterPath)
	resp, err := http.Get(url)
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
