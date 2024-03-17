package users

import "errors"

var (
	ErrEmailRequired        = errors.New("email is required")
	ErrPasswordRequired     = errors.New("password is required")
	ErrPasswordConfirmation = errors.New("password_confirmation does not match password")
)
