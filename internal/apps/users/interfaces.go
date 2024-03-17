package users

import (
	"context"
	"net/http"
)

type IService interface {
	CreateUser(w http.ResponseWriter, r *http.Request)
}

type ICases interface {
	CreateUser(ctx context.Context, user *User) error
}

type IRepository interface {
	CreateUser(ctx context.Context, user *User) error
}
