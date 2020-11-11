package unit_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"pkg.dsb.dev/unit"
)

func TestParseDataUnit(t *testing.T) {
	t.Parallel()

	tt := []struct {
		Name         string
		Input        string
		Expected     unit.DataUnit
		ExpectsError bool
	}{
		{
			Name:     "It should parse bytes",
			Input:    "5B",
			Expected: 5 * unit.Byte,
		},
		{
			Name:     "It should parse kilobytes",
			Input:    "10kB",
			Expected: 10 * unit.Kilobyte,
		},
		{
			Name:     "It should parse megabytes",
			Input:    "15MB",
			Expected: 15 * unit.Megabyte,
		},
		{
			Name:     "It should parse gigabytes",
			Input:    "20GB",
			Expected: 20 * unit.Gigabyte,
		},
		{
			Name:     "It should parse petabytes",
			Input:    "25PB",
			Expected: 25 * unit.Petabyte,
		},
		{
			Name:         "It should return an error for an invalid string",
			Input:        "25PB12",
			ExpectsError: true,
		},
		{
			Name:         "It should return an error for an invalid unit",
			Input:        "25BTC",
			ExpectsError: true,
		},
		{
			Name:     "It should handle simplified units",
			Input:    "1000B",
			Expected: unit.Kilobyte,
		},
	}

	for _, tc := range tt {
		t.Run(tc.Name, func(t *testing.T) {
			result, err := unit.ParseDataUnit(tc.Input)
			if tc.ExpectsError {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, tc.Expected, result)
		})
	}
}

func TestDataUnit_String(t *testing.T) {
	t.Parallel()

	tt := []struct {
		Name     string
		Unit     unit.DataUnit
		Expected string
	}{
		{
			Name:     "It should print bytes",
			Unit:     5 * unit.Byte,
			Expected: "5B",
		},
		{
			Name:     "It should print kilobytes",
			Unit:     5 * unit.Kilobyte,
			Expected: "5kB",
		},
		{
			Name:     "It should print megabytes",
			Unit:     5 * unit.Megabyte,
			Expected: "5MB",
		},
		{
			Name:     "It should print gigabytes",
			Unit:     5 * unit.Gigabyte,
			Expected: "5GB",
		},
		{
			Name:     "It should print terabytes",
			Unit:     5 * unit.Terabyte,
			Expected: "5TB",
		},
		{
			Name:     "It should print petabytes",
			Unit:     5 * unit.Petabyte,
			Expected: "5PB",
		},
		{
			Name:     "It should perform simplification",
			Unit:     2000 * unit.Kilobyte,
			Expected: "2MB",
		},
	}

	for _, tc := range tt {
		t.Run(tc.Name, func(t *testing.T) {
			result := tc.Unit.String()

			assert.Equal(t, tc.Expected, result)
		})
	}
}
