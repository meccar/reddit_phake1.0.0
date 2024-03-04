package api

import (
	"bytes"
	"fmt"
	"github.com/rs/zerolog/log"
	"net/http"
)

// "encoding/json"

func APIRequest(apiURL string, payload []byte) (*http.Response, error) {
	// Make a POST request to the API
	resp, err := http.Post(apiURL, "application/json", bytes.NewBuffer(payload))
	fmt.Printf("\n APIRequest apiURL %v\n", apiURL)

	if err != nil {
		return nil, fmt.Errorf("error making API request: %s", err)
	}
	fmt.Printf("\n APIRequest resp %v\n", resp)
	fmt.Printf("\n APIRequest err %v\n", err)

	defer resp.Body.Close()

	// Check the response status
	switch resp.StatusCode {
	case http.StatusOK:
		// Handle 200 (OK)
		log.Info().Msgf("\n API request succeeded. Response: %d\n", resp.StatusCode)
		return resp, nil
	case http.StatusBadRequest:
		// Handle 400 (Bad Request)
		log.Info().Msgf("\n API request returned 400 (Bad Request). Response: %d\n", resp.StatusCode)
		return resp, nil
	case http.StatusUnauthorized:
		// Handle 401 (Unauthorized)
		log.Info().Msgf("\n API request returned 401 (Unauthorized). Response: %d\n", resp.StatusCode)
		return resp, nil
	case http.StatusForbidden:
		// Handle 403 (Forbidden)
		log.Info().Msgf("\n API request returned 403 (Forbidden). Response: %d\n", resp.StatusCode)
		return resp, nil
	case http.StatusNotFound:
		// Handle 404 (Not Found)
		log.Info().Msgf("\n API request returned 404 (Not Found). Response: %d\n", resp.StatusCode)
		return resp, nil
	case http.StatusMethodNotAllowed:
		// Handle 405 (Method Not Allowed)
		log.Info().Msgf("\n API request returned 405 (Method Not Allowed). Response: %d\n", resp.StatusCode)
		return resp, nil
	case http.StatusConflict:
		// Handle 409 (Conflict)
		log.Info().Msgf("\n API request returned 409 (Conflict). Response: %d\n", resp.StatusCode)
		return resp, nil
	case http.StatusInternalServerError:
		// Handle 500 (Internal Server Error)
		log.Info().Msgf("\n API request returned 500 (Internal Server Error). Response: %d\n", resp.StatusCode)
		return resp, nil
	case http.StatusServiceUnavailable:
		// Handle 503 (Service Unavailable)
		log.Info().Msgf("\n API request returned 503 (Service Unavailable). Response: %d\n", resp.StatusCode)
		return resp, nil
	default:
		return nil, fmt.Errorf("API request failed with status: %d", resp.StatusCode)
	}
}
