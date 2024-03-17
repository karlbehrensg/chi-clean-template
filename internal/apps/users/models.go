package users

import (
	"errors"
	"net/url"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID           uuid.UUID
	Email        string
	PasswordHash string
	Active       bool
}

func (u *User) MountRequestValues(values url.Values) ([]error, error) {
	var errs []error
	if values.Get("email") == "" {
		errs = append(errs, ErrEmailRequired)
	}

	if values.Get("password") == "" {
		errs = append(errs, ErrPasswordRequired)
	}

	if values.Get("password_confirmation") != values.Get("password") {
		errs = append(errs, ErrPasswordConfirmation)
	}

	if len(errs) > 0 {
		return errs, errors.New("invalid form")
	}

	u.Email = values.Get("email")
	u.PasswordHash = values.Get("password")
	u.Active = true

	return nil, nil
}

func (u *User) HashPassword() error {
	if u.PasswordHash == "" {
		return ErrPasswordRequired
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.PasswordHash), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u.PasswordHash = string(hashedPassword)
	return nil
}
