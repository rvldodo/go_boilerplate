package api

import (
	"context"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"gorm.io/gorm"

	"github.com/rvldodo/boilerplate/app/handlers/auth"
	"github.com/rvldodo/boilerplate/app/handlers/google"
	"github.com/rvldodo/boilerplate/app/handlers/user"
	"github.com/rvldodo/boilerplate/app/middlewares"
	"github.com/rvldodo/boilerplate/domain/services"
	"github.com/rvldodo/boilerplate/domain/store"
	"github.com/rvldodo/boilerplate/lib/log"
	"github.com/rvldodo/boilerplate/lib/redis"
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
	ctx := context.Background()
	rds, _ := redis.New()
	err := redis.Run(ctx, rds.Client)
	if err != nil {
		log.Errorf("Redis failed: %v", err)
		return err
	}

	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedHeaders:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowCredentials: false,
		MaxAge:           300,
	})

	router := mux.NewRouter()
	router.Use(corsMiddleware.Handler)

	subrouter := router.PathPrefix("/api/v1").Subrouter()

	// User
	userStore := store.NewUserStore(s.DB)
	userService := services.NewUserService(userStore)
	authHandler := auth.NewHandler(userService)
	authHandler.RegisterRoutes(subrouter)

	userhandler := user.NewHandler(userService, *userStore)
	userhandler.RegisterRoutes(subrouter)

	googleService := services.NewGoogleService(userStore)
	googleHandler := google.NewHandler(googleService, userService)
	googleHandler.RegisterRoutes(subrouter)

	middlewares.LogginMiddleware(router)

	log.Infof("Server listening on: %s", s.Addrs)
	return http.ListenAndServe(s.Addrs, router)
}
