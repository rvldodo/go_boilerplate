package api

import (
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"

	"github.com/rvldodo/boilerplate/app/middlewares"
	"github.com/rvldodo/boilerplate/lib/log"
)

type ServerAPI struct {
	Addrs string
	Store *gorm.DB
}

func NewAPI(addrs string, store *gorm.DB) *ServerAPI {
	return &ServerAPI{
		Addrs: addrs,
		Store: store,
	}
}

func (s *ServerAPI) Run() error {
	router := mux.NewRouter()

	middlewares.LogginMiddleware(router)

	log.Infof("Server listening on: %s", s.Addrs)
	return http.ListenAndServe(s.Addrs, router)
}
