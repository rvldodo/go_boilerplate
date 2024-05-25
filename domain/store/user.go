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

func (us *UserStore) FindListUsers(
	ctx context.Context,
	limit, offset int,
) ([]model.UserResponse, int, int, error) {
	var users []model.User
	var count int64

	query := db.DB.WithContext(ctx).Model(&model.User{})
	if query.Error != nil {
		log.Errorf("Error find user list: %v", query.Error)
		return []model.UserResponse{}, 0, 0, query.Error
	}
	query.Count(&count)
	query.Find(&users)

	return buildUserListResponse(users), int(count), limit, nil
}

func (us *UserStore) FindByEmail(ctx context.Context, email string) (model.UserResponse, error) {
	var user model.User
	err := db.DB.WithContext(ctx).Where("email = ?", email).Find(&user).Error
	if err != nil {
		log.Errorf("Error find user by email: %v", err)
		return model.UserResponse{}, err
	}
	return buildUserResponse(user), nil
}

func (us *UserStore) FindByEmailShowPassword(
	ctx context.Context,
	email string,
) (model.UserResponseWithPassword, error) {
	var user model.User
	err := db.DB.WithContext(ctx).Where("email = ?", email).Find(&user).Error
	if err != nil {
		log.Errorf("Error find user by email: %v", err)
		return model.UserResponseWithPassword{}, err
	}
	return buildUserResponseWithPassword(user), nil
}

func (us *UserStore) FindById(ctx context.Context, id uuid.UUID) (model.UserResponse, error) {
	var user model.User
	err := db.DB.WithContext(ctx).Where("id = ?", id).Find(&user).Error
	if err != nil {
		log.Errorf("Error find user by id: %v", err)
		return model.UserResponse{}, err
	}
	return buildUserResponse(user), nil
}

func (us *UserStore) Update(
	ctx context.Context,
	id uuid.UUID,
	user model.UserRequestUpdate,
) (model.UserResponse, error) {
	var existUser model.User
	res := db.DB.WithContext(ctx).Where("id = ?", id).Find(&existUser)
	if res.Error != nil {
		log.Errorf("Error find user by id: %v", res.Error)
		return model.UserResponse{}, res.Error
	}

	res = db.DB.WithContext(ctx).Model(&existUser).Updates(&user)
	if res.Error != nil {
		log.Errorf("Error updating user: %v", res.Error)
		return model.UserResponse{}, res.Error
	}
	return buildUserResponse(existUser), nil
}

func (us *UserStore) DeleteById(ctx context.Context, id uuid.UUID) error {
	res := db.DB.WithContext(ctx).Where("id = ?", id).Delete(model.User{})
	if res.Error != nil {
		log.Errorf("Error delete user by id: %v", res.Error)
		return res.Error
	}
	return nil
}

func buildUserResponse(u model.User) model.UserResponse {
	return model.UserResponse{
		ID:           u.ID,
		FirstName:    u.FirstName,
		LastName:     u.LastName,
		Email:        u.Email,
		UserGoogleID: u.UserGoogleID,
		CreatedAt:    u.CreatedAt,
		UpdatedAt:    u.UpdatedAt,
	}
}

func buildUserResponseWithPassword(u model.User) model.UserResponseWithPassword {
	return model.UserResponseWithPassword{
		ID:        u.ID,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Email:     u.Email,
		Password:  u.Password,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}

func buildUserListResponse(u []model.User) []model.UserResponse {
	var userResponses []model.UserResponse
	for _, user := range u {
		userResponse := model.UserResponse{
			ID:        user.ID,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Email:     user.Email,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		}
		userResponses = append(userResponses, userResponse)
	}

	return userResponses
}
