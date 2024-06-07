package app

import (
	"go-http/internal/config"
	"go-http/internal/router"
)

func Start() {
	cfg := config.LoadEnv()
	log := config.LoadLogger(cfg.Env)

	log.Debug("123")

	// db

	r := router.Router{}
	r.InitRouters()
	r.LoadMiddleware()
	r.LoadRouters()
	r.ListenServer()
}
