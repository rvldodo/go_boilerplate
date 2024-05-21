package auth

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"

	"github.com/rvldodo/boilerplate/domain/model"
	"github.com/rvldodo/boilerplate/domain/services"
	"github.com/rvldodo/boilerplate/lib/log"
	"github.com/rvldodo/boilerplate/utils"
)

type Handler struct {
	service services.UserService
}

func NewHandler(service services.UserService) *Handler {
	return &Handler{service}
}

func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {
	var user model.UserRequest
	if err := utils.ParseJSON(r, &user); err != nil {
		log.Error(err)
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(user); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, errors)
		return
	}

	us, err := h.service.FindUserByEmail(r.Context(), user.Email)
	if err == nil && us.ID != uuid.Nil {
		log.Errorf("Email %s already registered", user.Email)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("Email already registered"))
		return
	}

	u, err := h.service.CreateUser(r.Context(), user)
	if err != nil {
		log.Error(err)
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, u)
}
