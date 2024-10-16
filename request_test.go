package openweather

import (
	"net/url"
	"testing"
)

func Test_encodedUrlValues(t *testing.T) {
	v := url.Values{}
	v.Set("lat", "44.10")
	v.Set("lon", "22.10")

	have := encodedUrlValues(v)

	want := "?lat=44.10&lon=22.10"

	if have != want {
		t.Errorf("have: %s, want: %s", have, want)
	}
}

func Test_requestUrl(t *testing.T) {
	v := url.Values{}
	v.Set("lat", "44.10")
	v.Set("lon", "22.10")

	have := requestUrl("weather", v)

	want := baseUrl + "/data/2.5/weather?lat=44.10&lon=22.10"

	if have != want {
		t.Errorf("have: %s, want: %s", have, want)
	}
}
