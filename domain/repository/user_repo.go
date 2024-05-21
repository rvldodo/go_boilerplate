package repository

import (
	"context"

	"github.com/google/uuid"

	"github.com/rvldodo/boilerplate/domain/model"
)

type UserRepository interface {
	Create(ctx context.Context, user model.User) error
	FindByEmail(ctx context.Context, email string) (model.UserResponse, error)
	FindById(ctx context.Context, id uuid.UUID) (model.UserResponse, error)
	Update(ctx context.Context, id uuid.UUID, user model.User) (model.UserResponse, error)
	DeleteById(ctx context.Context, id uuid.UUID) error
}
