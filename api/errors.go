package api

import (
    "errors"
)

var (
    ErrUnimplemented = errors.New("Unimplemented.")
    ErrInvalidItem = errors.New("Invalid Item")
)
