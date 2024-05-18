package base26

import "errors"

var (
	ErrInvalidBase26Char  = errors.New("invalid base26 character")
	ErrInvalidInputLength = errors.New("input string length must be in 1~15")
	ErrTooLargeForUint64  = errors.New("input string is too large for uint64")
)

type Base26Error struct {
	Input string
	Err   error
}

func (e *Base26Error) Error() string {
	return "input=" + e.Input + " : " + e.Err.Error()
}

func (e *Base26Error) Unwrap() error {
	return e.Err
}
