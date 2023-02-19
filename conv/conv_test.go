package conv

import (
	"math"
	"testing"
)

//	Mal for testfunksjoner
//	Du skal skrive alle funksjonene basert på denne malen
//	For alle konverteringsfunksjonene (tilsammen 6)
//	kan du bruke malen som den er (du må selvsagt endre
//	funksjonsnavn og testverdier)
//
// /
func withinTolerance(a, b, error float64) bool {
	// Først sjekk om tallene er nøyaktig like
	if a == b {
		return true
	}

	difference := math.Abs(a - b)

	// Siden vi skal dele med b, må vi sjekke om den er 0
	// Hvis b er 0, returner avgjørelsen om d er mindre enn feilmarginen
	// som vi aksepterer
	if b == 0 {
		return difference < error
	}
	return (difference / math.Abs(b)) < error
}
func TestFarhenheitToCelsius(t *testing.T) {
	type test struct {
		input float64
		want  float64
	}

	tests := []test{
		{input: 134, want: 56.67},
	}

	for _, tc := range tests {
		got := FarhenheitToCelsius(tc.input)
		if !withinTolerance(tc.want, got, 1e-4) {
			t.Errorf("expected: %v, got: %v", tc.want, got)
		}
	}
}
func TestFahrenheitToKevlin(t *testing.T) {
	type test struct {
		input float64
		want  float64
	}

	tests := []test{
		{input: 134, want: 329.82},
	}
	for _, tc := range tests {
		got := FarhenheitToKevlin(tc.input)

		if !withinTolerance(tc.want, got, 1e-4) {
			t.Errorf("expected: %.18f, got: %.18f", tc.want, got)
		}
	}
}
func TestKelvinToCelsius(t *testing.T) {
	type test struct {
		input float64
		want  float64
	}

	tests := []test{
		{input: 329.82, want: 56.67},
	}

	for _, tc := range tests {
		got := KelvinToCelsius(tc.input)

		if !withinTolerance(tc.want, got, 1e-12) {
			t.Errorf("expected: %.18f, got: %.18f", tc.want, got)
		}
	}
}
func TestKelvinToFahrenheit(t *testing.T) {
	type test struct {
		input float64
		want  float64
	}

	tests := []test{
		{input: 329.82, want: 134},
	}

	for _, tc := range tests {
		got := KelvinToFarhenheit(tc.input)

		if !withinTolerance(tc.want, got, 1e-4) {
			t.Errorf("expected: %.18f, got: %.18f", tc.want, got)
		}
	}
}
func TestCelsiusToFarhenheit(t *testing.T) {
	type test struct {
		input float64
		want  float64
	}

	tests := []test{
		{input: 56.67, want: 134},
	}

	for _, tc := range tests {
		got := CelsiusToFarhenheit(tc.input)

		if !withinTolerance(tc.want, got, 1e-4) {
			t.Errorf("expected: %.18f, got: %.18f", tc.want, got)
		}
	}
}

func TestCelsiusToKelvin(t *testing.T) {
	type test struct {
		input float64
		want  float64
	}

	tests := []test{
		{input: 56.67, want: 329.82},
	}

	for _, tc := range tests {
		got := CelsiusToKelvin(tc.input)

		if !withinTolerance(tc.want, got, 1e-4) {
			t.Errorf("expected: %.18f, got: %.18f", tc.want, got)
		}
	}
}
