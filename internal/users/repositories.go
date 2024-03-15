package users

import (
	"context"

	"github.com/karlbehrensg/chi-clean-template/utils"
)

type repository struct{}

func newRepository() IRepository {
	return &repository{}
}

func (r *repository) CreateUser(ctx context.Context, user *User) error {
	ctx = utils.NewTrace(ctx)
	utils.InfoTrace(ctx, "Creating user in db")
	return nil
}
