package routers

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
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

func TestRedirectURL(t *testing.T) {
	r := RedirectRouter()

	type addTest struct {
		key  string
		link string
	}

	var tests []addTest

	for _, record := range records {
		tests = append(tests, addTest{
			key:  record.Key,
			link: record.Value,
		})
	}

	tests = append(tests, addTest{
		key:  "abcdefgh",
		link: "localhost:3000",
	})

	for _, test := range tests {
		url := fmt.Sprintf("/%s", test.key)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", url, nil)
		r.ServeHTTP(w, req)

		want := true
		got := strings.Contains(w.Body.String(), test.link)

		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}

		got = strings.Contains("302,303", fmt.Sprintf("%v", w.Code))

		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}

	}
}

func TestShortenURL(t *testing.T) {
	r := ApiRouter()

	type addTest struct {
		key   string
		link  string
		title string
	}

	var tests []addTest

	for _, record := range records {
		tests = append(tests, addTest{
			key:   record.Key,
			link:  record.Value,
			title: record.Title,
		})
	}

	for _, test := range tests {

		body := strings.NewReader(fmt.Sprintf(`{"url": "%s"}`, test.link))

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/shorten", body)
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		r.ServeHTTP(w, req)

		want := true
		got := strings.Contains(w.Body.String(), test.title)

		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}

		got = strings.Contains("200", fmt.Sprintf("%v", w.Code))

		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/shorten", strings.NewReader(`{"url": "notgonnawork"}`))
	r.ServeHTTP(w, req)

	got := w.Code
	want := 500

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestGetDetails(t *testing.T) {
	r := ApiRouter()

	type addTest struct {
		key   string
		link  string
		title string
	}

	var tests []addTest

	for _, record := range records {
		tests = append(tests, addTest{
			key:   record.Key,
			link:  record.Value,
			title: record.Title,
		})
	}

	for _, test := range tests {
		url := fmt.Sprintf("/%s", test.key)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", url, nil)
		r.ServeHTTP(w, req)

		got := strings.Contains("200", fmt.Sprintf("%v", w.Code))
		want := true

		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}

		got = strings.Contains(w.Body.String(), test.link)

		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}

		got = strings.Contains(w.Body.String(), test.title)

		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}

	}

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/abcdefsdg", nil)
	r.ServeHTTP(w, req)

	got := w.Code
	want := 404

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}
