package repository

import (
	"context"
	"database/sql"
	"errors"

	model "github.com/erdedigital/gokit"
	"github.com/go-kit/kit/log"
)

var RepoErr = errors.New("Unable to handle Repo Request")

type repo struct {
	db     *sql.DB
	logger log.Logger
}

func NewRepo(db *sql.DB, logger log.Logger) Repository {
	return &repo{
		db:     db,
		logger: log.With(logger, "repo", "sql"),
	}
}

func (repo *repo) CreateUser(ctx context.Context, user model.User) error {
	sql := `
		INSERT INTO users (id, email, password)
		VALUES ($1, $2, $3)`

	if user.Email == "" || user.Password == "" {
		return RepoErr
	}

	_, err := repo.db.ExecContext(ctx, sql, user.ID, user.Email, user.Password)
	if err != nil {
		return err
	}
	return nil
}

func (repo *repo) GetUser(ctx context.Context, id string) (*model.User, error) {
	var u model.User
	err := repo.db.QueryRow("SELECT email, password FROM users WHERE id=$1", id).Scan(&u.Email, &u.Password)
	if err != nil {
		return nil, RepoErr
	}

	return &u, nil
}
