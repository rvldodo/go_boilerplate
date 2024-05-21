package user

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/google/uuid"
	"github.com/gorilla/mux"

	"github.com/rvldodo/boilerplate/domain/model"
	"github.com/rvldodo/boilerplate/domain/services"
	"github.com/rvldodo/boilerplate/domain/store"
	"github.com/rvldodo/boilerplate/lib/log"
	"github.com/rvldodo/boilerplate/utils"
)

type Handler struct {
	service services.UserService
	store   store.UserStore
}

func NewHandler(service services.UserService, store store.UserStore) *Handler {
	return &Handler{service, store}
}

func (h *Handler) handleListUsers(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	limit, err := strconv.Atoi(query.Get("limit"))
	if err != nil {
		limit = 5
	}

	offset, err := strconv.Atoi(query.Get("page"))
	if err != nil {
		offset = 0
	}

	res, count, err := h.service.FindListUsers(r.Context(), limit, offset)
	if err != nil {
		log.Errorf("Failed get user list: %v", err)
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, map[string]interface{}{"users": res, "totalPage": count})
}

func (h *Handler) handleDeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	str := vars["userID"]

	userID, _ := uuid.Parse(str)

	err := h.service.DeleteUserById(r.Context(), userID)
	if err != nil {
		log.Errorf("Error user not found by id: %v", err)
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, fmt.Sprintf("Successfully delete user by id %v", userID))
}

func (h *Handler) handleUpdateUser(w http.ResponseWriter, r *http.Request) {
	var user model.UserRequestUpdate

	vars := mux.Vars(r)
	str := vars["userID"]
	userID, _ := uuid.Parse(str)

	if err := utils.ParseJSON(r, &user); err != nil {
		log.Errorf("Error parse user: %v", err)
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	u, err := h.service.UpdateUser(r.Context(), userID, user)
	if err != nil {
		log.Errorf("Error in update user: %v", err)
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, u)
}
