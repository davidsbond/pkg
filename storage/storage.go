// Package storage contains utilities for persistence methods used by the application.
package storage

import (
	"time"

	"github.com/jackc/pgtype"
)

// FromTextArray converts an instance of pgtype.TextArray to a string array.
func FromTextArray(ta pgtype.TextArray) []string {
	out := make([]string, len(ta.Elements))
	for i, elem := range ta.Elements {
		out[i] = elem.String
	}
	return out
}

// ToTextArray converts a given string slice to a pgtype.TextArray instance.
// Will panic if Set returns an error.
func ToTextArray(arr []string) pgtype.TextArray {
	tags := pgtype.TextArray{}
	if err := tags.Set(arr); err != nil {
		panic(err)
	}
	return tags
}

// ToInterval converts a given time.Duration to a pgtype.Interval instance.
// Will panic if Set returns an error.
func ToInterval(dur time.Duration) pgtype.Interval {
	interval := pgtype.Interval{}
	if err := interval.Set(dur); err != nil {
		panic(err)
	}
	return interval
}

// FromInterval converts a given pgtype.Interval instance into a time.Duration
// instance. It assumes the interval is set using microseconds, as that is how
// `ToInterval` will create them.
func FromInterval(interval pgtype.Interval) time.Duration {
	const conversion = 1000
	return time.Duration(interval.Microseconds * conversion)
}
