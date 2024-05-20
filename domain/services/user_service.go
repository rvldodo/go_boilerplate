package services

import (
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

func (s *UserService) CreateUser(user model.UserRequest) (model.UserResponse, error) {
	u := model.NewUser(&user)
	err := s.repo.Create(user)
	if err != nil {
		log.Errorf("Failed create user: %v", err)
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

func (s *UserService) FindUserByID(id uuid.UUID) (model.UserResponse, error) {
	u, err := s.repo.FindById(id)
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

func (s *UserService) FindUserByEmail(email string) (model.UserResponse, error) {
	u, err := s.repo.FindByEmail(email)
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
