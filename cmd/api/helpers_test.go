package main

import (
	"net/http"
	"testing"
)

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

		got, _ := GetHtmlTitle(res.Body)
		want := test.title

		if got != want {
			t.Errorf("got %q, wanted %q", got, want)
		}
	}

}