package users

import (
	"context"

	"github.com/karlbehrensg/chi-clean-template/utils"
)

type cases struct {
	repo IRepository
}

func newCases(repo IRepository) ICases {
	return &cases{
		repo,
	}
}

func (c *cases) CreateUser(ctx context.Context, user *User) error {
	ctx = utils.NewTrace(ctx)

	utils.InfoTrace(ctx, "Creating user")

	c.repo.CreateUser(ctx, user)

	return nil
}
