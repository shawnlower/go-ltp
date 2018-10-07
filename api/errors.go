package api

import (
	"errors"
)

var (
	ErrUnimplemented = errors.New("Unimplemented.")
	ErrInvalidItem   = errors.New("Invalid Item")
	ErrInvalidUri   = errors.New("Specified URI is invalid")
)
