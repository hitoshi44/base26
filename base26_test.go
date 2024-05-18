package base26

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var uintTests = []struct {
	name string
	u    uint64
	s    string
}{
	{"0", 0, "a"},
	{"1", 1, "b"},
	{"2", 2, "c"},
	{"3", 3, "d"},
	{"25", 25, "z"},
	{"26", 26, "ba"},
	{"27", 27, "bb"},
	{"28", 28, "bc"},
	{"52", 52, "ca"},
	{"78", 78, "da"},
	{"650", 650, "za"},
	{"675", 675, "zz"},
	{"676", 676, "baa"},
	{"677", 677, "bab"},
	{"123123123", 123123123, "kjleut"},
	{"max uint64", maxUint64, maxBase26String},
}

func TestEncodeUint(t *testing.T) {
	for _, tt := range uintTests {
		t.Run(tt.name, func(t *testing.T) {
			got := EncodeUint(tt.u)
			assert.Equal(t, tt.s, got)
		})
	}
}

func TestDecodeUint(t *testing.T) {
	for _, tt := range uintTests {
		t.Run(tt.name, func(t *testing.T) {
			got := MustDecodeUint(tt.s)
			assert.Equal(t, tt.u, got)
		})
	}
}

func TestDecodeUintError(t *testing.T) {
	var err error
	// too big for uint64
	_, err = DecodeUint("hlhxczmxsyumqq")
	assert.ErrorIs(t, err, ErrTooLargeForUint64)
	// empty string
	_, err = DecodeUint("")
	assert.ErrorIs(t, err, ErrInvalidInputLength)
	// too long string
	_, err = DecodeUint("hlhxczmxsyumqpp")
	assert.ErrorIs(t, err, ErrInvalidInputLength)
	// invalid character
	_, err = DecodeUint("a!")
	assert.ErrorIs(t, err, ErrInvalidBase26Char)
	_, err = DecodeUint("👺")
	assert.ErrorIs(t, err, ErrInvalidBase26Char)
}
