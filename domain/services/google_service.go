package services

import (
	"context"

	"github.com/rvldodo/boilerplate/domain/model"
	"github.com/rvldodo/boilerplate/domain/repository"
	"github.com/rvldodo/boilerplate/domain/store"
	"github.com/rvldodo/boilerplate/lib/log"
)

type GoogleService struct {
	repo repository.UserRepository
}

func NewGoogleService(repo *store.UserStore) GoogleService {
	return GoogleService{repo}
}

func (gs *GoogleService) CreateUser(
	ctx context.Context,
	user model.UserGoogle,
) (model.UserResponse, error) {
	u := model.UserRequest{
		FirstName:    user.GivenName,
		LastName:     user.FamilyName,
		Email:        user.Email,
		UserGoogleID: user.ID,
	}

	// ue, err := gs.repo.FindByEmail(ctx, user.Email)
	// if err == nil && ue.ID != uuid.Nil {
	// 	log.Error("User already registered")
	// 	return model.UserResponse{}, fmt.Errorf("User already registered")
	// }

	us := model.NewUser(&u)
	err := gs.repo.Create(ctx, us)
	if err != nil {
		log.Errorf("Error create user from google: %v", err)
		return model.UserResponse{}, err
	}

	return model.UserResponse{
		ID:           us.ID,
		FirstName:    us.FirstName,
		LastName:     us.LastName,
		Email:        us.Email,
		UserGoogleID: us.UserGoogleID,
		CreatedAt:    us.CreatedAt,
		UpdatedAt:    &us.CreatedAt,
	}, nil
}
