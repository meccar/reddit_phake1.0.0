package util

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetImageBase64(imagePath string) string {
	// Make an HTTP GET request to fetch the image data
	resp, err := http.Get(imagePath)
	if err != nil {
		fmt.Println("Error fetching image:", err)
		return ""
	}
	defer resp.Body.Close()

	// Read the response body
	imageBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading image data:", err)
		return ""
	}

	// Encode the image to base64
	base64Encoded := base64.StdEncoding.EncodeToString(imageBytes)
	return base64Encoded
}
