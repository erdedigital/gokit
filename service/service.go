package service

import (
	"context"

	"github.com/erdedigital/gokit/model"
)

type Service interface {
	CreateUser(ctx context.Context, email string, password string) (string, error)
	GetUser(ctx context.Context, id string) (*model.User, error)
}
