package error

import (
	"net/http"
	"strings"
)

// RegisterStatusHandler handles the error status and message
func StatusHandler(msg error) int {
	errorMsg := msg.Error()
	errorMessages := strings.Split(errorMsg, ", ")


	// Iterate through the errors and return the first non-zero status and error message
	for _, err := range errorMessages {
		status := GetHTTPStatus(err)
		if status != 0 {
			return status
		}
	}
	// If no specific error found, return default status and message
	return GetHTTPStatus(errorMsg)
}

// GetHTTPStatus returns the appropriate HTTP status code based on the error message
func GetHTTPStatus(errorMessage string) int {
	// Check the error message and return the corresponding HTTP status code
	switch errorMessage {
	case "Invalid Username", "Invalid Password", "Email is already registered", "Invalid Name", "Invalid Email", "Invalid Phone":
		return http.StatusConflict // 409 Conflict
	case "Password must be at least X characters long.", "Password must contain uppercase, lowercase letters, numbers, and symbols.", "Password does not meet security requirements.":
		return http.StatusBadRequest // 400 Bad Request
	case "Username is Incorrect", "Password is Incorrect", "token is unauthorized", "token is expired":
		return http.StatusUnauthorized // 401 Status Unauthorized
	default:
		return http.StatusInternalServerError // Default 500 Status Internal Server Error
	}
}
