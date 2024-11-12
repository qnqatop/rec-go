package db

import (
	"encoding/json"
	"fmt"
	"strconv"
)

// JSONNumber is a drop-in replacement for old (before go1.14) json.Number that accepts empty string as valid value.
// Problem: https://github.com/golang/go/issues/37308
type JSONNumber string

var (
	_ json.Marshaler   = JSONNumber("")
	_ json.Unmarshaler = new(JSONNumber)
)

var _ = json.Number("")

func (n JSONNumber) String() string { return string(n) }

// Float64 returns the number as a float64.
func (n JSONNumber) Float64() (float64, error) {
	return strconv.ParseFloat(string(n), 64)
}

// Int64 returns the number as an int64.
func (n JSONNumber) Int64() (int64, error) {
	return strconv.ParseInt(string(n), 10, 64)
}

func (n JSONNumber) MarshalJSON() ([]byte, error) {
	if n.String() == "" {
		return []byte("null"), nil
	}
	return []byte(n), nil
}

func (n *JSONNumber) UnmarshalJSON(b []byte) error {
	if len(b) == 0 {
		return nil
	}
	v := string(b)

	if b[0] == '"' { // unquote when string
		var err error
		v, err = strconv.Unquote(v)
		if err != nil {
			return err
		}
	}

	if !isValidNumber(v) {
		return fmt.Errorf("invalid number literal %q", v)
	}
	*n = JSONNumber(v)
	return nil
}

// isValidNumber reports whether s is a valid JSON number literal.
// copy-pasted from go source code of json package.
//
//nolint:all
func isValidNumber(s string) bool {
	// This function implements the JSON numbers grammar.
	// See https://tools.ietf.org/html/rfc7159#section-6
	// and https://www.json.org/img/number.png

	if len(s) == 0 {
		return true
	}

	// Optional -
	if s[0] == '-' {
		s = s[1:]
		if s == "" {
			return false
		}
	}

	// Digits
	switch {
	default:
		return false

	case s[0] == '0':
		s = s[1:]

	case '1' <= s[0] && s[0] <= '9':
		s = s[1:]
		for len(s) > 0 && '0' <= s[0] && s[0] <= '9' {
			s = s[1:]
		}
	}

	// . followed by 1 or more digits.
	if len(s) >= 2 && s[0] == '.' && '0' <= s[1] && s[1] <= '9' {
		s = s[2:]
		for len(s) > 0 && '0' <= s[0] && s[0] <= '9' {
			s = s[1:]
		}
	}

	// e or E followed by an optional - or + and
	// 1 or more digits.
	if len(s) >= 2 && (s[0] == 'e' || s[0] == 'E') {
		s = s[1:]
		if s[0] == '+' || s[0] == '-' {
			s = s[1:]
			if s == "" {
				return false
			}
		}
		for len(s) > 0 && '0' <= s[0] && s[0] <= '9' {
			s = s[1:]
		}
	}

	// Make sure we are at the end.
	return s == ""
}
