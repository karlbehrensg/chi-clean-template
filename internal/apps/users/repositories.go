package users

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/karlbehrensg/chi-clean-template/utils"
	_ "github.com/lib/pq"
)

type repository struct {
	db *sql.DB
}

func newRepository(db *sql.DB) IRepository {
	return &repository{
		db: db,
	}
}

func (r *repository) CreateUser(ctx context.Context, user *User) error {
	var id uuid.UUID

	ctx = utils.NewTrace(ctx)
	utils.InfoTrace(ctx, "Creating user in db")

	err := r.db.QueryRowContext(
		ctx,
		`INSERT INTO users (email, password_hash) VALUES ($1, $2) RETURNING id`,
		user.Email, user.PasswordHash,
	).Scan(&id)
	if err != nil {
		logUser := user
		logUser.PasswordHash = "[MASKED]"
		utils.ErrorTrace(ctx, "Error creating user in db: ", err, *logUser)
		return err
	}

	user.ID = id

	utils.InfoTrace(ctx, fmt.Sprintf("User %s created successfully", user.Email))

	return nil
}
