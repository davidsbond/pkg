// Package random contains utilities for generating random things.
package random

import (
	"crypto/rand"
	"io"
	"math/big"
	"strings"
)

// Int generates a random integer between 0 and max. Panics if it fails
// to read the RNG.
func Int(max int) int {
	i, err := rand.Int(rand.Reader, big.NewInt(int64(max)))
	if err != nil {
		panic(err)
	}

	return int(i.Int64())
}

// String returns a random string of a given length.
func String(length int) string {
	k := make([]byte, length)
	if _, err := io.ReadFull(rand.Reader, k); err != nil {
		panic(err)
	}
	return string(k)
}

// CapitalisedAlphaNumeric returns a random string consisting of alphanumeric characters
// where all letters are capitalised. Expects a length greater than zero.
func CapitalisedAlphaNumeric(length int) string {
	const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	var builder strings.Builder
	for i := 0; i < length; i++ {
		j := Int(len(charset))
		char := rune(charset[j])
		builder.WriteRune(char)
	}

	return builder.String()
}
