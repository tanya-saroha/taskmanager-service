package api

import (
	"errors"
)

const (
	DuplicateRecordpg          = "23505"
	ErrorForeignKeyViolationpg = "23503"
)

var (
	//common
	ErrJSONFormat = errors.New("Invalid request")
	ErrEmptyValue = errors.New("Value not present")
	ErrBadUser    = errors.New("Provided User is invalid")
	ErrAuthUser   = errors.New("You are not an authorized user to perform this action")

	//middleware
	ErrAuthNotPresent    = errors.New("Authentication header required")
	ErrUserIDNotPresent  = errors.New("User ID is not present")
	ErrExpiryNotPresent  = errors.New("Expiration time is not present")
	ErrBadUserID         = errors.New("Invalid User ID")
	ErrProfileNotPresent = errors.New("User Profile is not present")
	//ErrTokenNotPresent   = errors.New("token not present")
)
