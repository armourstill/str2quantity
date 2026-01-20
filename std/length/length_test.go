package length

import (
	"math"
	"testing"
)

func TestParseLength(t *testing.T) {
	tests := []struct {
		input string
		want  float64 // in meters
	}{
		// SI Units
		{"1m", 1.0},
		{"1.5m", 1.5},
		{"1km", 1000.0},
		{"100cm", 1.0},
		{"1000mm", 1.0},
		{"1µm", 1e-6},
		{"1um", 1e-6},
		{"1nm", 1e-9},

		// Multipart
		{"1m 50cm", 1.5},
		{"1km 500m", 1500.0},
	}

	epsilon := 1e-9

	for _, tt := range tests {
		got, err := ParseLength(tt.input)
		if err != nil {
			t.Errorf("ParseLength(%q) unexpected error: %v", tt.input, err)
			continue
		}
		if math.Abs(got-tt.want) > epsilon {
			t.Errorf("ParseLength(%q) = %v, want %v", tt.input, got, tt.want)
		}
	}
}

func TestParseLength_Errors(t *testing.T) {
	invalidInputs := []string{
		"1kg",    // Wrong unit
		"hello",  // Garbage
		"",       // Empty
		"1.1.1m", // Bad number
	}

	for _, input := range invalidInputs {
		_, err := ParseLength(input)
		if err == nil {
			t.Errorf("ParseLength(%q) expected error, got nil", input)
		}
	}
}
