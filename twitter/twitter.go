package twitter

import "errors"

var (
	ErrBadCredentials     = errors.New("incorrect email/password combination")
	ErrNotFound           = errors.New("not found")
	ErrValidation         = errors.New("validation error")
	ErrInvalidAccessToken = errors.New("invalid access token")
	ErrNoUserIDInContext  = errors.New("no user id in context")
	ErrGenAccessToken     = errors.New("generate access token error")
	ErrUnauthenticated    = errors.New("unauthenticated")
)
