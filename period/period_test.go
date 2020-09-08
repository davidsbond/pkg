package period_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"pkg.dsb.dev/period"
)

func TestParse(t *testing.T) {
	t.Parallel()

	tt := []struct {
		Name     string
		Input    string
		Expected period.Period
		Ok       bool
	}{
		{
			Name:  "It should handle an invalid input",
			Input: "hello",
		},
		{
			Name:     "It should parse a day",
			Input:    "day",
			Expected: period.Day,
			Ok:       true,
		},
		{
			Name:     "It should parse a week",
			Input:    "week",
			Expected: period.Week,
			Ok:       true,
		},
		{
			Name:     "It should parse a month",
			Input:    "month",
			Expected: period.Month,
			Ok:       true,
		},
		{
			Name:     "It should parse a year",
			Input:    "year",
			Expected: period.Year,
			Ok:       true,
		},
	}

	for _, tc := range tt {
		t.Run(tc.Name, func(t *testing.T) {
			actual, ok := period.Parse(tc.Input)

			assert.Equal(t, tc.Expected, actual)
			assert.Equal(t, tc.Ok, ok)
		})
	}
}
