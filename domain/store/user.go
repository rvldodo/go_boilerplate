package store

import (
	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/rvldodo/boilerplate/domain/model"
)

type UserStore struct {
	db *gorm.DB
}

func NewUserStore(db *gorm.DB) *UserStore {
	return &UserStore{db: db}
}

func (us *UserStore) Create(user model.UserRequest) error {
	return nil
}

func (us *UserStore) FindByEmail(email string) (model.UserResponse, error) {
	return model.UserResponse{}, nil
}

func (us *UserStore) FindById(id uuid.UUID) (model.UserResponse, error) {
	return model.UserResponse{}, nil
}

func (us *UserStore) Update(id uuid.UUID, user model.UserRequest) (model.UserResponse, error) {
	return model.UserResponse{}, nil
}

func (us *UserStore) DeleteById(id uuid.UUID) error {
	return nil
}
