package user

import (
	"github.com/gorilla/mux"

	"github.com/rvldodo/boilerplate/app/middlewares"
)

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/users", middlewares.AuthWithJWT(h.handleListUsers, h.store))
	router.HandleFunc("/users/{userID}", middlewares.AuthWithJWT(h.handleDeleteUser, h.store)).
		Methods("DELETE")
}
