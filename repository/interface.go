package repository

import (
	"context"

	model "github.com/erdedigital/gokit"
)

type Repository interface {
	CreateUser(ctx context.Context, user model.User) error
	GetUser(ctx context.Context, id string) (*model.User, error)
}
