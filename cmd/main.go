package main

import (
	"orders/config"
	"orders/internal"
)

func main() {
	cfg := config.New()
	server := internal.New(cfg.Env, cfg.Port)
	server.Run()
}
