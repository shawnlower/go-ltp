package server


import (
	"fmt"

	"google.golang.org/grpc/codes"
)

// Error is the structured error returned from the server
type Error struct {
	// Code is the canonical error code for describing the nature of a
	// particular error.
	Code codes.Code
	// Desc explains more details of the error.
	Desc string
}

func (e *Error) Error() string {
	if e == nil {
		return fmt.Sprintf("ltpd: OK")
	}
	return fmt.Sprintf("ltpd: code = %q, desc = %q", e.Code, e.Desc)
}
