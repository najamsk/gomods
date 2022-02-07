// Package str has miscellaneous string functions
// Uses ascii package for lower/upper case
package str

import "github.com/apmckinlay/gsuneido/util/ascii"

// Capitalized returns true is the string starts with A-Z, otherwise false
func Capitalized(s string) bool {
	return len(s) >= 1 && ascii.IsUpper(s[0])
}

// Capitalize returns the string with the first letter converted from a-z to A-Z
func Capitalize(s string) string {
	if len(s) > 0 && ascii.IsLower(s[0]) {
		s = string(ascii.ToUpper(s[0])) + s[1:]
	}
	return s
}

// UnCapitalize returns the string with the first letter converted from A-Z to a-z
func UnCapitalize(s string) string {
	if len(s) > 0 && ascii.IsUpper(s[0]) {
		s = string(ascii.ToLower(s[0])) + s[1:]
	}
	return s
}

// IndexFunc returns the index of the first byte that the func returns true for
// else -1
func IndexFunc(s string, f func(byte) bool) int {
	for i, c := range []byte(s) {
		if f(c) {
			return i
		}
	}
	return -1
}

// Dup is intended to make a copy of a string
// so we don't hold a reference to a large source and prevent garbage collection
func Dup(s string) string {
	s = " " + s
	return s[1:]
}

// IndexNotAny returns the index of the first byte not in chars, else -1
func IndexNotAny(s, chars string) int {
	nc := len(chars)
loop:
	for i := 0; i < len(s); i++ {
		c := s[i]
		for j := 0; j < nc; j++ {
			if c == chars[j] {
				continue loop
			}
		}
		return i
	}
	return -1
}

// LastIndexNotAny returns the last index of the first byte not in chars, else -1
func LastIndexNotAny(s, chars string) int {
	nc := len(chars)
loop:
	for i := len(s) - 1; i >= 0; i-- {
		c := s[i]
		for j := 0; j < nc; j++ {
			if c == chars[j] {
				continue loop
			}
		}
		return i
	}
	return -1
}

// Doesc returns the next byte, interpreting escape sequences
func Doesc(s string, i int) (byte, int) {
	c := s[i]
	if c != '\\' || i+1 >= len(s) {
		return c, i
	}
	c = s[i+1]
	switch c {
	case 'n':
		return '\n', i + 1
	case 't':
		return '\t', i + 1
	case 'r':
		return '\r', i + 1
	case '\\', '"', '\'':
		return c, i + 1
	case 'x':
		if i+2 < len(s) {
			dig1 := ascii.Digit(s[i+1], 16)
			dig2 := ascii.Digit(s[i+2], 16)
			if dig1 != -1 && dig2 != -1 {
				return byte(16*dig1 + dig2), i + 3
			}
		}
	}
	return '\\', i
}
