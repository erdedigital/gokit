package service

import (
	"context"

	model "github.com/erdedigital/gokit"
)

type Service interface {
	CreateUser(ctx context.Context, email string, password string) (string, error)
	GetUser(ctx context.Context, id string) (*model.User, error)
}
