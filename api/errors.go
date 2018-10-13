package api

import (
    "fmt"
)

var (
    ErrUnimplemented = fmt.Errorf("Unimplemented.")
    ErrInvalidItem = fmt.Errorf("Invalid item")
)

type ErrInvalidAuthMethod struct {
    Method string
}

func (e ErrInvalidAuthMethod) Error() string {
    return fmt.Sprintf("Invalid authentication method `%s'. Valid methods are 'grpc'", e.Method)
}

type ErrInvalidScheme struct {
    Scheme string
}

func (e ErrInvalidScheme) Error() string {
    return fmt.Sprintf("Invalid authentication scheme `%s'. Valid schemes are 'mutual-tls', and 'insecure'", e.Scheme)
}

type ErrInvalidUri struct {
    Uri string
}

func (e ErrInvalidUri) Error() string {
    return fmt.Sprintf("Specified URI: `%s' is invalid", e.Uri)
}


