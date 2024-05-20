package repository

import (
	"github.com/google/uuid"

	"github.com/rvldodo/boilerplate/domain/model"
)

type UserRepository interface {
	Create(user model.UserRequest) error
	FindByEmail(email string) (model.UserResponse, error)
	FindById(id uuid.UUID) (model.UserResponse, error)
	Update(id uuid.UUID, user model.UserRequest) (model.UserResponse, error)
	DeleteById(id uuid.UUID) error
}
