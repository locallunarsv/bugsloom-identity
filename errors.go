package identity

import "errors"

var (
	ErrInvalidID = errors.New("invalid identity")
	ErrEmptyID   = errors.New("empty identity")
)