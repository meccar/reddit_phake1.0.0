package util

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type ResponseData struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
}

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	// Create ResponseData instance
	responseData := ResponseData{
		Status: status,
		Data:   v,
	}

	// Encode JSON data
	jsonData, err := json.MarshalIndent(responseData, "", "  ")
	if err != nil {
		return fmt.Errorf("error encoding JSON data: %w", err)
	}

	// Set content type header
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	// Write JSON data to the response writer
	_, err = w.Write(jsonData)
	if err != nil {
		return fmt.Errorf("error writing response body: %w", err)
	}

	return nil
}

func ReadJSON(url string, v interface{}) error {
	// Make a GET request to the specified URL
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("error making GET request: %w", err)
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("error reading response body: %w", err)
	}

	// Parse the JSON data
	if err := json.Unmarshal(body, v); err != nil {
		return fmt.Errorf("error parsing JSON data: %w", err)
	}

	return nil
}
