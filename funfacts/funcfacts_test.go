package funfacts

import (
	"reflect"
	"testing"
)

func TestGetFunFacts(t *testing.T) {
	// Set up FunFacts struct
	ff := FunFacts{
		Sun:   []string{"Temperatur på ytre lag av Solen 5778°K", "Temperatur i Solens kjerne 15000000°C"},
		Luna:  []string{"Temperatur på Månens overflate om natten -183°C", "Temperatur på Månens overflate om dagen 106°C"},
		Terra: []string{"Høyeste temperatur målt på Jordens overflate 56.7°C", "Laveste temperatur målt på Jordens overflate -89.4°C"},
	}

	tests := []struct {
		input string
		want  []string
	}{
		{input: "sun", want: []string{"Temperatur på ytre lag av Solen 5778°K", "Temperatur i Solens kjerne 15000000°C"}},
		{input: "luna", want: []string{"Temperatur på Månens overflate om natten -183°C", "Temperatur på Månens overflate om dagen 106°C"}},
		{input: "terra", want: []string{"Høyeste temperatur målt på Jordens overflate 56.7°C", "Laveste temperatur målt på Jordens overflate -89.4°C"}},
		{input: "invalid", want: []string{"Sorry, I don't have any fun facts about that."}},
	}

	for _, tc := range tests {
		got := ff.GetFunFacts(tc.input)
		if !reflect.DeepEqual(tc.want, got) {
			t.Errorf("GetFunFacts(%q) = %v, want %v", tc.input, got, tc.want)
		}
	}
}
