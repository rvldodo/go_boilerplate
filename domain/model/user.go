package model

import (
	"time"

	"github.com/google/uuid"

	"github.com/rvldodo/boilerplate/lib/bcrypt"
)

type User struct {
	ID        uuid.UUID  `json:"id,omitempty"`
	FirstName string     `json:"first_name,omitempty"`
	LastName  string     `json:"last_name,omitempty"`
	Email     string     `json:"email,omitempty"`
	Password  string     `json:"password,omitempty"`
	CreatedAt time.Time  `json:"created_at,omitempty" gorm:"column:created_at;type:timestamp;not null;autoCreateTime"`
	UpdatedAt *time.Time `json:"updated_at,omitempty" gorm:"column:updated_at;type:timestamp ON UPDATE CURRENT_TIMESTAMP;null;autoUpdateTime"`
}

type UserRequest struct {
	FirstName string `json:"first_name,omitempty" validate:"required"`
	LastName  string `json:"last_name,omitempty"  validate:"required"`
	Email     string `json:"email,omitempty"      validate:"required,email"`
	Password  string `json:"password,omitempty"   validate:"required,min=6"`
}

type UserResponse struct {
	ID        uuid.UUID
	FirstName string
	LastName  string
	Email     string
	CreatedAt time.Time
	UpdatedAt *time.Time
}

func NewUser(user *UserRequest) User {
	hash, _ := bcrypt.HashedPassword(user.Password)

	return User{
		ID:        uuid.New(),
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Password:  hash,
	}
}
