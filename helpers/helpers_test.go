package helpers

import (
	"net/http"
	"testing"
	"time"

	"github.com/berabulut/kapsule/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var records = map[string]*models.ShortURL{
	"_Ng8-6_8T": {
		ID:        primitive.NewObjectID(),
		Key:       "_Ng8-6_8T",
		Value:     "https://github.com/public-apis/public-apis",
		CreatedAt: time.Now().Unix(),
		Clicks:    0,
		Visits:    []models.Visit{},
		Title:     "GitHub - public-apis/public-apis: A collective list of free APIs",
	},
}

func TestGetCountryCode(t *testing.T) {

	type addTest struct {
		IP, CountryCode string
	}

	var tests = []addTest{
		{"1.1.1.1", "AUS"},
		{"4.1.1.2", "USA"},
	}

	for _, test := range tests {
		got, _ := GetCountryCode(test.IP)
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

		got, _ := GetHtmlTitle(res.Body)
		want := test.title

		if got != want {
			t.Errorf("got %q, wanted %q", got, want)
		}
	}

}

func TestHandleClick(t *testing.T) {
	type addTest struct {
		key    string
		clicks int
	}

	var tests []addTest

	for _, record := range records {
		tests = append(tests, addTest{
			key:    record.Key,
			clicks: record.Clicks,
		})
	}

	for _, test := range tests {
		HandleClick(records[test.key], models.UserAgent{}, "", "")

		got := records[test.key].Clicks
		want := test.clicks + 1

		if got != want {
			t.Errorf("got %d, wanted %d", got, want)
		}
	}
}
