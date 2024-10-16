package openweather

import (
	"fmt"
	"net/url"
)

var baseUrl = "https://api.openweathermap.org"

type Request struct {
	Lat   float64
	Lon   float64
	Lang  string
	Units string
}

func (r Request) url(path string, token string) string {
	v := url.Values{}

	v.Set("lat", fmt.Sprintf("%.2f", r.Lat))

	v.Set("lon", fmt.Sprintf("%.2f", r.Lon))

	v.Set("appid", token)

	setUrlValue("lang", r.Lang, "en", &v)

	setUrlValue("units", r.Units, "metric", &v)

	return requestUrl(path, v)
}

func setUrlValue(key string, value string, defaultVal string, v *url.Values) {
	if value == "" {
		v.Set(key, defaultVal)
	} else {
		v.Set(key, value)
	}
}

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

type requestBuilder interface {
	url(path string, token string) string
}
