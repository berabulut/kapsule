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
	r := RedirectRouter(records)

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
