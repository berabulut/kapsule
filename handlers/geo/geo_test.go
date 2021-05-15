package geo

import "testing"

func TestGetCountryCode(t *testing.T) {

	got, _ := GetCountryCode("1.1.1.1")
	want := "AUS"

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
