package endpoint

import (
	"context"

	repo "github.com/erdedigital/gokit/repository"
	service "github.com/erdedigital/gokit/service"
	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	CreateUser endpoint.Endpoint
	GetUser    endpoint.Endpoint
}

func MakeEndpoints(s service.Service) Endpoints {
	return Endpoints{
		CreateUser: makeCreateUserEndpoint(s),
		GetUser:    makeGetUserEndpoint(s),
	}
}

func makeCreateUserEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(repo.CreateUserRequest)
		ok, err := s.CreateUser(ctx, req.Email, req.Password)
		return repo.CreateUserResponse{Ok: ok}, err
	}
}

func makeGetUserEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(repo.GetUserRequest)
		resp, err := s.GetUser(ctx, req.Id)

		return repo.GetUserResponse{
			Email:    resp.Email,
			Password: resp.Password,
		}, err
	}
}
