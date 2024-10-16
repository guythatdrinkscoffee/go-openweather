package openweather

import (
	"compress/gzip"
	"encoding/json"
	"io"
	"net/http"
)

func validateHttpResponse(response *http.Response) error {
	if response.StatusCode == 200 {
		return nil
	}

	errRes := &ErrorResponse{}
	if err := bind(response, &errRes); err != nil {
		return err
	}

	return &APIError{
		Response:      response,
		ErrorResponse: errRes,
	}
}

func bind(response *http.Response, into interface{}) error {
	body, err := readResponse(response)
	if err != nil {
		return err
	}

	bytes, err := io.ReadAll(body)
	if err != nil {
		return err
	}

	if len(bytes) < 1 || bytes == nil {
		return nil
	}

	return json.Unmarshal(bytes, &into)
}

func readResponse(response *http.Response) (io.Reader, error) {
	h := response.Header.Get("Content-Encoding")

	if len(h) < 1 {
		return response.Body, nil
	}

	return gzip.NewReader(response.Body)
}
