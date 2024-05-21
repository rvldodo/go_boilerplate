package user

import (
	"net/http"
	"strconv"

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
