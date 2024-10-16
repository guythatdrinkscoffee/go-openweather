package openweather

import (
	"fmt"
	"net/http"
)

type ErrorResponse struct {
	StatusCode int      `json:"cod"`
	Message    string   `json:"message"`
	Parameters []string `json:"parameters"`
}

type APIError struct {
	Response      *http.Response
	ErrorResponse *ErrorResponse
}

func (a *APIError) Error() string {
	return fmt.Sprintf("http: %d, error: %s", a.Response.StatusCode, a.ErrorResponse.Message)
}
