package openweather

import (
	"fmt"
	"net/url"
)

var baseUrl = "https://api.openweathermap.org"

func requestUrl(path string, values url.Values) string {
	return baseUrl + fmt.Sprintf("/data/2.5/%s", path) + encodedUrlValues(values)
}

func encodedUrlValues(values url.Values) string {
	q := values.Encode()

	if q == "" {
		return q
	}

	return "?" + q
}
