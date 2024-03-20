package util

import (
	"io"
	"net/http"
)

func GetImageBinary(imagePath string) []byte {
	// Make an HTTP GET request to fetch the image data
	resp, err := http.Get(imagePath)
	if err != nil {
		return nil
	}
	defer resp.Body.Close()

	// Read the response body
	imageBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil
	}

	return imageBytes
}
