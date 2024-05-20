package main

import (
	"github.com/rvldodo/boilerplate/app/api"
	"github.com/rvldodo/boilerplate/config"
	"github.com/rvldodo/boilerplate/db"
	"github.com/rvldodo/boilerplate/lib/log"
)

func main() {
	server := api.NewAPI(config.Envs.Addrs, db.DB)
	if err := server.Run(); err != nil {
		log.Errorf("Failed running server: %v", err)
	}
}
