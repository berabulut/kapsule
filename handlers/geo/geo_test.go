package geo

import "testing"

type addTest struct {
	ip, country_code string
}

var tests = []addTest{
	{"1.1.1.1", "AUS"},
	{"4.1.1.2", "USA"},
}

func TestGetCountryCode(t *testing.T) {

	for _, test := range tests {
		got, _ := GetCountryCode(test.ip)
		want := test.country_code

		if got != want {
			t.Errorf("got %q, wanted %q", got, want)
		}
	}

}
