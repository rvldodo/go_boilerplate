package services

import (
	"context"

	"github.com/google/uuid"

	"github.com/rvldodo/boilerplate/domain/model"
	"github.com/rvldodo/boilerplate/domain/repository"
	"github.com/rvldodo/boilerplate/domain/store"
	"github.com/rvldodo/boilerplate/lib/log"
)

type UserService struct {
	repo repository.UserRepository
}

func NewUserService(repo *store.UserStore) UserService {
	return UserService{repo}
}

func (s *UserService) CreateUser(
	ctx context.Context,
	user model.UserRequest,
) (model.UserResponse, error) {
	u := model.NewUser(&user)
	err := s.repo.Create(ctx, u)
	if err != nil {
		log.Errorf("Failed create user: %v", err)
		return model.UserResponse{}, err
	}

	return model.UserResponse{
		ID:        u.ID,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Email:     u.Email,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}, nil
}

func (s *UserService) FindUserByID(ctx context.Context, id uuid.UUID) (model.UserResponse, error) {
	u, err := s.repo.FindById(ctx, id)
	if err != nil {
		log.Errorf("Failed find user by id: %v", err)
		return model.UserResponse{}, err
	}

	return model.UserResponse{
		ID:        u.ID,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Email:     u.Email,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}, nil
}

func (s *UserService) FindUserByEmail(
	ctx context.Context,
	email string,
) (model.UserResponse, error) {
	u, err := s.repo.FindByEmail(ctx, email)
	if err != nil {
		log.Errorf("Failed find user by email: %v", err)
		return model.UserResponse{}, err
	}

	return model.UserResponse{
		ID:        u.ID,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Email:     u.Email,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}, nil
}

func (s *UserService) FindUserByEmailShowPassword(
	ctx context.Context,
	email string,
) (model.UserResponseWithPassword, error) {
	u, err := s.repo.FindByEmailShowPassword(ctx, email)
	if err != nil {
		log.Errorf("Failed find user by email: %v", err)
		return model.UserResponseWithPassword{}, err
	}

	return model.UserResponseWithPassword{
		ID:        u.ID,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Email:     u.Email,
		Password:  u.Password,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}, nil
}

func (s *UserService) FindListUsers(
	ctx context.Context,
	limit, offset int,
) ([]model.UserResponse, int, error) {
	res, count, _, err := s.repo.FindListUsers(ctx, limit, offset)
	if err != nil {
		log.Errorf("Failed get list users: %v", err)
		return []model.UserResponse{}, 0, err
	}

	return res, int(count / limit), nil
}
