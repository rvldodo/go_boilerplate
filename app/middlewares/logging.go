package middlewares

import (
	"github.com/gorilla/mux"

	"github.com/rvldodo/boilerplate/lib/log"
)

func LogginMiddleware(router *mux.Router) *mux.Router {
	router.Use(log.LoggingMiddleware)
	return router
}
