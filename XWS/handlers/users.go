package handlers

import (
	"fmt"
	followProtos "follows_service/protos/follows"
	"log"
	postProtos "posts_service/protos/posts"
	protos "users_service/protos/user"
	"xws_proj/data"
)

// KeyUser is a key used for the User object in the context
type KeyUser struct{}

// Users handler for getting and updating users
type Users struct {
	l  *log.Logger
	v  *data.Validation
	uc protos.UserClient
	pc postProtos.PostsClient
	fc followProtos.FollowsClient
}

// NewUsers returns a new products handler with the given logger
func NewUsers(l *log.Logger, v *data.Validation, uc protos.UserClient, pc postProtos.PostsClient, fc followProtos.FollowsClient) *Users {
	return &Users{l, v, uc, pc, fc}
}

// ErrInvalidUserPath is an error message when the user path is not valid
var ErrInvalidUserPath = fmt.Errorf("invalid Path, path should be /users/[id]")

// GenericError is a generic error message returned by a server
type GenericError struct {
	Message string `json:"message"`
}

// ValidationError is a collection of validation error messages
type ValidationError struct {
	Messages []string `json:"messages"`
}
