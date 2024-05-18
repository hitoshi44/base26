package base26

import "strings"

const digits26 = "abcdefghijklmnopqrstuvwxyz"

const maxBase26String = "hlhxczmxsyumqp"
const maxUint64 = 18446744073709551615

// EncodeUint encodes a uint64 to a base26 string.
func EncodeUint(u uint64) string {

	var result [64]byte
	var i = 64

	for u >= 26 {
		i--
		q, r := divideBy26(u)
		u = q
		result[i] = digits26[r]
	}
	i-- // u < 26
	result[i] = digits26[uint(u)]

	return string(result[i:])
}

// MustDecodeUint decodes a base26 string (lower case) to a uint64.
// It panics if the input string is invalid or too large for uint64.
// Use this function if you are sure that the input string is valid.
func MustDecodeUint(s string) uint64 {
	var result uint64
	for i := 0; i < len(s); i++ {
		result = result*26 + uint64(s[i]-'a')
	}
	return result
}

// DecodeUint decodes a base26 string to a uint64.
func DecodeUint(s string) (uint64, error) {

	// input length check
	if len(s) < 1 || 14 < len(s) {
		return 0, &Base26Error{s, ErrInvalidInputLength}
	}

	lower := strings.ToLower(s)

	// check only base26 characters
	for _, r := range s {
		if r < 'a' || r > 'z' {
			return 0, &Base26Error{lower, ErrInvalidBase26Char}
		}
	}
	// check the upper bound
	if lower > "hlhxczmxsyumqp" {
		return 0, &Base26Error{lower, ErrTooLargeForUint64}
	}

	return MustDecodeUint(lower), nil
}

// Avoid using r = a%b in addition to q = a/b
// reason: 64bit division and modulo operations
// see go standard library strconv/itoa.go.
func divideBy26(u uint64) (q, r uint64) {
	q = u / 26
	r = u - q*26
	return
}
