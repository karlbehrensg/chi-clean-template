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

	utils.InfoTrace(ctx, "Hashing user password")
	if err := user.HashPassword(); err != nil {
		logUser := user
		logUser.PasswordHash = "[MASKED]"
		utils.ErrorTrace(ctx, "Error validating user fields: ", err, *logUser)
		return err
	}

	utils.InfoTrace(ctx, "Creating user")
	c.repo.CreateUser(ctx, user)

	return nil
}
