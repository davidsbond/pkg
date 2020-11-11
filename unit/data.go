// Package unit contains types and methods for dealing with units of various types. Currently supports metric
// data units.
package unit

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

type (
	// The DataUnit type contains a metric representation of a unit of data, ie bytes, kilobytes etc.
	DataUnit int
)

// Supported DataUnit values.
const (
	Byte     = DataUnit(1)
	Kilobyte = Byte * 1000
	Megabyte = Kilobyte * 1000
	Gigabyte = Megabyte * 1000
	Terabyte = Gigabyte * 1000
	Petabyte = Terabyte * 1000
)

var (
	// ErrInvalidDataUnit is the error given when parsing a data unit that contains invalid
	// characters.
	ErrInvalidDataUnit = errors.New("invalid data unit")

	// ErrUnknownDataUnit is the error given when parsing a data unit and the unit string does not match
	// any known units.
	ErrUnknownDataUnit = errors.New("unknown data unit")

	dataUnitValues = map[string]DataUnit{
		"B":  Byte,
		"kB": Kilobyte,
		"MB": Megabyte,
		"GB": Gigabyte,
		"TB": Terabyte,
		"PB": Petabyte,
	}

	dataUnitNames = map[DataUnit]string{
		Byte:     "B",
		Kilobyte: "kB",
		Megabyte: "MB",
		Gigabyte: "GB",
		Terabyte: "TB",
		Petabyte: "PB",
	}
)

func (du DataUnit) String() string {
	var u DataUnit
	switch {
	case du > Petabyte:
		u = Petabyte
	case du > Terabyte:
		u = Terabyte
	case du > Gigabyte:
		u = Gigabyte
	case du > Megabyte:
		u = Megabyte
	case du > Kilobyte:
		u = Kilobyte
	default:
		u = Byte
	}

	return fmt.Sprint(int(du/u), dataUnitNames[u])
}

// ParseDataUnit attempts to return a DataUnit corresponding to the provided string. Current values are supported
// from byte to petabyte:
//
// 10B: 10 bytes
// 10kB: 10 kilobytes
// 20MB: 20 megabytes
// 5GB: 5 gigabytes
// 10TB: 10 terabytes
// 1PB: 1 petabyte.
func ParseDataUnit(str string) (DataUnit, error) {
	number := strings.Builder{}
	unit := strings.Builder{}

	for _, r := range str {
		switch {
		case unicode.IsNumber(r):
			// We shouldn't have any unit if we're parsing a number.
			if unit.Len() > 0 {
				return 0, fmt.Errorf("%w: %s", ErrInvalidDataUnit, str)
			}

			number.WriteRune(r)
		case unicode.IsLetter(r):
			unit.WriteRune(r)
		default:
			return 0, fmt.Errorf("%w: %s", ErrInvalidDataUnit, str)
		}
	}

	unitStr := unit.String()
	multiplier, ok := dataUnitValues[unitStr]
	if !ok {
		return 0, fmt.Errorf("%w: %s", ErrUnknownDataUnit, unitStr)
	}

	amount, err := strconv.ParseInt(number.String(), 10, 64)
	if err != nil {
		return 0, err
	}

	return DataUnit(amount) * multiplier, nil
}
