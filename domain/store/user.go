package store

import (
	"context"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/rvldodo/boilerplate/db"
	"github.com/rvldodo/boilerplate/domain/model"
	"github.com/rvldodo/boilerplate/lib/log"
)

type UserStore struct {
	db *gorm.DB
}

func NewUserStore(db *gorm.DB) *UserStore {
	return &UserStore{db: db}
}

func (us *UserStore) Create(ctx context.Context, user model.User) error {
	res := db.DB.WithContext(ctx).Create(&user)
	if res.Error != nil {
		log.Errorf("Error user repo: %v", res.Error)
		return res.Error
	}
	return nil
}

func (us *UserStore) FindByEmail(ctx context.Context, email string) (model.UserResponse, error) {
	var user model.User
	err := db.DB.WithContext(ctx).Where("email = ?", email).Find(&user).Error
	if err != nil {
		log.Errorf("Error find user by email: %v", err)
		return model.UserResponse{}, err
	}
	return model.UserResponse{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, nil
}

func (us *UserStore) FindById(ctx context.Context, id uuid.UUID) (model.UserResponse, error) {
	return model.UserResponse{}, nil
}

func (us *UserStore) Update(
	ctx context.Context,
	id uuid.UUID,
	user model.User,
) (model.UserResponse, error) {
	return model.UserResponse{}, nil
}

func (us *UserStore) DeleteById(ctx context.Context, id uuid.UUID) error {
	return nil
}
