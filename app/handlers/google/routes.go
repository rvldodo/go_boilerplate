package google

import "github.com/gorilla/mux"

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/google_login", h.handleGoogle).Methods("GET")
	router.HandleFunc("/google_callback", h.handleGoogleCallback).Methods("GET")
}
