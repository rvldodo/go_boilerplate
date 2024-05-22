package google

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/rvldodo/boilerplate/config"
	"github.com/rvldodo/boilerplate/domain/model"
	"github.com/rvldodo/boilerplate/domain/services"
	"github.com/rvldodo/boilerplate/lib/log"
	"github.com/rvldodo/boilerplate/utils"
)

type Handler struct {
	service services.GoogleService
}

func NewHandler(service services.GoogleService) *Handler {
	return &Handler{service}
}

func (h *Handler) handleGoogle(w http.ResponseWriter, r *http.Request) {
	url := config.AppConfig.GoogleLoginConfig.AuthCodeURL(config.Envs.GoogleRandomState)

	http.Redirect(w, r, url, http.StatusSeeOther)
}

func (h *Handler) handleGoogleCallback(w http.ResponseWriter, r *http.Request) {
	state := r.URL.Query().Get("state")
	if state != config.Envs.GoogleRandomState {
		log.Error("Invalid google state")
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("Invalid google state"))
		return
	}

	code := r.URL.Query().Get("code")
	token, err := config.GoogleInit.Exchange(r.Context(), code)
	if err != nil {
		log.Errorf("Invalid code: %v", err)
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	resp, err := http.Get(
		"https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken,
	)
	if err != nil {
		log.Errorf("Error fetch user data: %v", err)
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	defer resp.Body.Close()

	var userInfo model.UserGoogle
	err = json.NewDecoder(resp.Body).Decode(&userInfo)
	if err != nil {
		log.Errorf("Error parse json: %v", err)
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	us, err := h.service.CreateUser(r.Context(), userInfo)
	if err != nil {
		log.Errorf("Error create user from google: %v", err)
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, us)
}
