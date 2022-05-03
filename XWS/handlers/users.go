package handlers

import (
	"fmt"
	"log"
	"xws_proj/data"
)

// KeyProduct is a key used for the Product object in the context
type KeyUser struct{}

// Products handler for getting and updating products
type Users struct {
	l *log.Logger
	v *data.Validation
}

// NewProducts returns a new products handler with the given logger
func NewUsers(l *log.Logger, v *data.Validation) *Users {
	return &Users{l, v}
}

// ErrInvalidUserPath is an error message when the user path is not valid
var ErrInvalidUserPath = fmt.Errorf("Invalid Path, path should be /users/[id]")

// GenericError is a generic error message returned by a server
type GenericError struct {
	Message string `json:"message"`
}

// ValidationError is a collection of validation error messages
type ValidationError struct {
	Messages []string `json:"messages"`
}
