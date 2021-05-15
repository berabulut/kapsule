package handlers

import (
	"net/http"
	"testing"
)

func TestGetCountryCode(t *testing.T) {

	type addTest struct {
		IP, CountryCode string
	}

	var tests = []addTest{
		{"1.1.1.1", "AUS"},
		{"4.1.1.2", "USA"},
	}

	for _, test := range tests {
		got, _ := getCountryCode(test.IP)
		want := test.CountryCode

		if got != want {
			t.Errorf("got %q, wanted %q", got, want)
		}
	}

}

func TestGetHtmlTitle(t *testing.T) {

	type addTest struct {
		url, title string
	}

	tests := []addTest{
		{"https://github.com/berabulut", "berabulut (Hüseyin Bera Bulut) · GitHub"},
		{"https://www.digitalocean.com/", "DigitalOcean – The developer cloud"},
	}

	for _, test := range tests {
		res, err := http.Get(test.url)
		if err != nil {
			t.Errorf("%s", err)
		}

		got, _ := getHtmlTitle(res.Body)
		want := test.title

		if got != want {
			t.Errorf("got %q, wanted %q", got, want)
		}
	}

}
