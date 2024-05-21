package api

import (
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"

	"github.com/rvldodo/boilerplate/app/handlers/auth"
	"github.com/rvldodo/boilerplate/app/handlers/user"
	"github.com/rvldodo/boilerplate/app/middlewares"
	"github.com/rvldodo/boilerplate/domain/services"
	"github.com/rvldodo/boilerplate/domain/store"
	"github.com/rvldodo/boilerplate/lib/log"
)

type ServerAPI struct {
	Addrs string
	DB    *gorm.DB
}

func NewAPI(addrs string, store *gorm.DB) *ServerAPI {
	return &ServerAPI{
		Addrs: addrs,
		DB:    store,
	}
}

func (s *ServerAPI) Run() error {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	// User
	userStore := store.NewUserStore(s.DB)
	userService := services.NewUserService(userStore)
	authHandler := auth.NewHandler(userService)
	authHandler.RegisterRoutes(subrouter)

	userhandler := user.NewHandler(userService, *userStore)
	userhandler.RegisterRoutes(subrouter)

	middlewares.LogginMiddleware(router)

	log.Infof("Server listening on: %s", s.Addrs)
	return http.ListenAndServe(s.Addrs, router)
}
