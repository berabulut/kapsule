package mongo

import (
	"testing"
)

var keys = []string{"DMWuGZ_KT", "EvCLfT_8T"}

func TestGetRecord(t *testing.T) {

	for _, key := range keys {

		record, err := GetRecord(key)

		if err != nil {
			t.Error(err)
		}

		got := record.Key
		want := key

		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}

	}

	_, err := GetRecord("EvCLfT_8Ts")

	if err == nil {
		t.Error("This shouldn't have passed!")
	}

}
