package length

import (
	"errors"

	"github.com/armourstill/str2quantity/parser"
	"github.com/armourstill/str2quantity/unit"
)

// System is the shared unit system for Length operations.
var System *unit.System

func init() {
	// Initialize system for Length strings.
	// We allow multipart (e.g., "1m 50cm") and stick to case-sensitivity for SI units.
	System = unit.NewSystem(unit.SystemConfig{
		AllowMultiPart:  true,
		CaseInsensitive: false,
	})

	// Base Unit: Meter (m)
	System.Add("m", 1.0, unit.DimLength)

	// SI Prefixes for Meter
	prefixes := []struct {
		sym string
		val float64
	}{
		{"n", 1e-9}, // nanometer
		{"u", 1e-6}, // micrometer
		{"µ", 1e-6}, // micrometer symbol
		{"m", 1e-3}, // millimeter
		{"c", 1e-2}, // centimeter
		{"k", 1e3},  // kilometer
	}

	for _, p := range prefixes {
		System.AddPrefix(p.sym, p.val, "m")
	}
}

// ParseLength parses a length string into meters (float64).
func ParseLength(s string) (float64, error) {
	val, dim, err := parser.Parse[float64](s, System)
	if err != nil {
		return 0, err
	}

	if !dim.Equals(unit.DimLength) {
		return 0, errors.New("parsed quantity is not a length")
	}

	return val, nil
}
